package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/hudl/vorpal/models"
)

func AllCharacters() (models.Characters, error) {
	var characters models.Characters
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting all characters")
		err := tx.Select(&characters, `SELECT * FROM characters`)
		if err != nil {
			return fmt.Errorf("Failed to get all characters: %v", err)
		}

		log.Debug("Successfully got all characters")
		return nil
	})
	return characters, err
}

func GetCharacter(id int) (*models.Character, error) {
	var character models.Character
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting character with id=%d", id)
		err := tx.Get(&character, `SELECT * FROM characters WHERE id=$1`, id)
		if err != nil {
			return fmt.Errorf("Failed to get character with id=%d: %v", id, err)
		}

		log.Debug("Successfully got character=%+v", character)
		return nil
	})
	return &character, err
}

func GetCharacterByName(name string) (*models.Character, error) {
	var character models.Character
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting character=%q", name)
		err := tx.Get(&character, `SELECT * FROM characters WHERE name=$1`, name)
		if err != nil {
			return fmt.Errorf("Failed to get character=%q: %v", name, err)
		}

		log.Debug("Successfully got character=%+v", character)
		return nil
	})
	return &character, err
}
