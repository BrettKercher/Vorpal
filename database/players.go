package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/hudl/vorpal/models"
)

type PlayerAllTimeStats struct {
	Player        *models.Player
	Kills         int `db:"kills"`
	Deaths        int `db:"deaths"`
	SelfDestructs int `db:"self_destructs"`
}

func AllPlayers() (models.Players, error) {
	var players models.Players
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting all players")
		err := tx.Select(&players, `SELECT * FROM players`)
		if err != nil {
			return fmt.Errorf("Failed to get all players: %v", err)
		}

		log.Debug("Successfully got all players")
		return nil
	})
	return players, err
}

func GetPlayer(id int) (*models.Player, error) {
	var player models.Player
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting player with id=%d", id)
		err := tx.Get(&player, `SELECT * FROM players WHERE id=$1`, id)
		if err != nil {
			return fmt.Errorf("Failed to get player with id=%d: %v", id, err)
		}

		log.Debug("Successfully got player=%+v", player)
		return nil
	})
	return &player, err
}

func GetPlayerByEmail(email string) (*models.Player, error) {
	var player models.Player
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting player with email=%q", email)
		err := tx.Get(&player, `SELECT * FROM players WHERE email=$1`, email)
		if err != nil {
			return fmt.Errorf("Failed to get player with email=%q: %v", email, err)
		}

		log.Debug("Successfully got player=%+v", player)
		return nil
	})
	return &player, err
}

func GetPlayerAllTimeStats(id int) (*PlayerAllTimeStats, error) {
	player, err := GetPlayer(id)
	if err != nil {
		return nil, err
	}

	stmt := `SELECT sum(kills) AS kills, sum(deaths) AS deaths, sum(self_destructs) AS self_destructs FROM player_stats WHERE player_id=$1`

	var stats PlayerAllTimeStats
	err = withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting all time stats or player=%q", player.Email)
		err := tx.Get(&stats, stmt, player.Id)
		if err != nil {
			return fmt.Errorf("Failed to get all time stats for player=%q: %v", player.Email, err)
		}

		log.Debug("Successfully got all time stats for player=%q", player.Email)
		return nil
	})

	return &stats, err
}

func CreatePlayer(player *models.Player) (*models.Player, error) {
	err := withTx(func(tx *sqlx.Tx) error {
		stmt := `INSERT INTO players (
			first_name,
			last_name,
			email,
			tag,
			is_disabled
		) VALUES ($1,$2,$3,$4,$5) RETURNING id`

		log.Debug("Creating player=%+v", player)
		var id int
		err := tx.Get(&id, stmt,
			player.FirstName,
			player.LastName,
			player.Email,
			player.Tag,
			player.IsDisabled)

		if err != nil {
			return fmt.Errorf("Failed to create player=%+v: %v", player, err)
		}

		// update the player id
		player.Id = id

		log.Debug("Successfully created player=%+v", player)
		return nil
	})
	return player, err
}

func UpdatePlayerName(id int, firstName string, lastName string) error {
	return withTx(func(tx *sqlx.Tx) error {
		log.Debug("Updating first_name=%q last_name=%q for player with id=%d",
			firstName, lastName, id)
		_, err := tx.Exec(`UPDATE players SET first_name=$1, last_name=$2 WHERE id=$3`,
			firstName, lastName, id)
		if err != nil {
			return fmt.Errorf("Failed to update first_name=%q last_name=%q for player with id=%d: %v",
				firstName, lastName, id, err)
		}

		log.Debug("Successfully updated first_name=%q last_name=%q for player with id=%d",
			firstName, lastName, id)
		return nil
	})
}

func UpdatePlayerTag(id int, tag string) error {
	return withTx(func(tx *sqlx.Tx) error {
		log.Debug("Updating tag=%q for player with id=%d", tag, id)
		_, err := tx.Exec(`UPDATE players SET tag=$1 WHERE id=$2`, tag, id)
		if err != nil {
			return fmt.Errorf("Failed to update tag=%q for player with id=%d: %v", tag, id, err)
		}

		log.Debug("Successfully updated tag=%q for player with id=%d", tag, id)
		return nil
	})
}

func DisabledPlayer(id int) error {
	return withTx(func(tx *sqlx.Tx) error {
		log.Debug("Disabling player with id=%d", id)
		_, err := tx.Exec(`UPDATE players SET is_disabled=true WHERE id=$1`, id)
		if err != nil {
			return fmt.Errorf("Failed to disable player with id=%d: %v", id, err)
		}

		log.Debug("Successfully disabled user with id=%d", id)
		return nil
	})
}
