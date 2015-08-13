package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/hudl/vorpal/models"
)

func AllStages() (models.Stages, error) {
	var stages models.Stages
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting all stages")
		err := tx.Select(&stages, `SELECT * FROM stages`)
		if err != nil {
			return fmt.Errorf("Failed to get all stages: %v", err)
		}

		log.Debug("Successfully got all stages")
		return nil
	})
	return stages, err
}

func GetStage(id int) (*models.Stage, error) {
	var stage models.Stage
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting stage with id=%d", id)
		err := tx.Get(&stage, `SELECT * FROM stages WHERE id=$1`, id)
		if err != nil {
			return fmt.Errorf("Failed to get stage with id=%d: %v", id, err)
		}

		log.Debug("Successfully got stage=%+v", stage)
		return nil
	})
	return &stage, err
}

func GetStageByName(name string) (*models.Stage, error) {
	var stage models.Stage
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting stage=%q", name)
		err := tx.Get(&stage, `SELECT * FROM stages WHERE name=$1`, name)
		if err != nil {
			return fmt.Errorf("Failed to get stage=%q: %v", name, err)
		}

		log.Debug("Successfully got stage=%+v", stage)
		return nil
	})
	return &stage, err
}
