package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	golog "github.com/op/go-logging"
)

const (
	serializable = "SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;"
)

///////////////////////////////////////////////////////////////////////////////
// Globals

var log = golog.MustGetLogger("database")
var utc, _ = time.LoadLocation("UTC")

var db *DB

///////////////////////////////////////////////////////////////////////////////
// DB

type DB struct {
	sqlx.DB
	MaxOpen int
}

// Opens a new connection to the database and saves the connection as the
// global db
func Open(dataSourceName string) error {
	log.Debug("Opening a new database connection")
	sqlxDB, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return fmt.Errorf("Failed to open a new database connection: %v", err)
	}

	log.Debug("Creating new database connection")
	db = &DB{DB: *sqlxDB}

	log.Debug("Setting the maximum open connections")
	db.SetMaxOpenConns(db.MaxOpen)

	return nil
}

func Close() error {
	return db.Close()
}

func (db *DB) Begin() (*sqlx.Tx, error) {
	log.Debug("Beginning new transaction")
	tx, err := db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("Failed to begin new transaction: %v", err)
	}

	log.Debug("Setting transaction isolation level to serializable")
	_, err = tx.Exec(serializable)
	if err != nil {
		return nil, fmt.Errorf("Failed to set transaction isolation level to serializable: %v", err)
	}

	return tx, nil
}

///////////////////////////////////////////////////////////////////////////////
// withTx

type withTxFunc func(tx *sqlx.Tx) error

func withTx(f withTxFunc) error {
	log.Debug("Getting transaction")
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Failed to get transaction: %v", err)
	}

	// call the argument function
	err = f(tx)
	if err != nil {
		return err
	}

	log.Debug("Committing transaction")
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Failed to commit transaction: %v", err)
	}

	// success
	return nil
}
