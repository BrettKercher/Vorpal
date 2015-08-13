package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/hudl/vorpal/models"
)

///////////////////////////////////////////////////////////////////////////////
// Intermediate Structs

// used to read in foreign keys for relations

type game struct {
	models.Game
	StageId int `db:"stage_id"`
}

type playerStat struct {
	models.PlayerStat
	GameId      int `db:"game_id"`
	PlayerId    int `db:"player_id"`
	CharacterId int `db:"character_id"`
}

///////////////////////////////////////////////////////////////////////////////
// Player Stats

func getPlayerStat(tx *sqlx.Tx, id int) (*models.PlayerStat, error) {
	var playerStat playerStat
	log.Debug("Getting player stat with id=%d", id)
	err := tx.Get(&playerStat, `SELECT * FROM player_stats WHERE id=$1`, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get player stat with id=%d: %v", id, err)
	}

	playerStat.Player, err = GetPlayer(playerStat.PlayerId)
	if err != nil {
		return nil, err
	}

	playerStat.Character, err = GetCharacter(playerStat.CharacterId)
	if err != nil {
		return nil, err
	}

	log.Debug("Successfully got player_stat=%+v", playerStat)
	return &playerStat.PlayerStat, nil
}

func getPlayerStatsForGame(tx *sqlx.Tx, gameId int) (models.PlayerStats, error) {
	var ids []int
	log.Debug("Getting player stats for game with id=%d", gameId)
	err := tx.Select(&ids, `SELECT id FROM player_stats WHERE game_id=$1`, gameId)
	if err != nil {
		return nil, fmt.Errorf("Failed to get player stats for game with id=%d: %v", gameId, err)
	}

	var playerStats models.PlayerStats
	for _, id := range ids {
		playerStat, err := getPlayerStat(tx, id)
		if err != nil {
			return nil, err
		}
		playerStats = append(playerStats, playerStat)
	}

	log.Debug("Successfully got player stats for game with id=%d", gameId)
	return playerStats, nil
}

func createPlayerStat(tx *sqlx.Tx, playerStat *models.PlayerStat, gameId int) (*models.PlayerStat, error) {
	stmt := `INSERT INTO player_stats (
		game_id,
		player_id,
		number,
		rank,
		character_id,
		is_random,
		kills,
		deaths,
		self_destructs
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id`

	var id int
	log.Debug("Creating player_stat=%+v", playerStat)
	err := tx.Get(&id, stmt,
		gameId,
		playerStat.Player.Id,
		playerStat.Number,
		playerStat.Rank,
		playerStat.Character.Id,
		playerStat.IsRandom,
		playerStat.Kills,
		playerStat.Deaths,
		playerStat.SelfDestructs)
	if err != nil {
		return nil, fmt.Errorf("Failed to create player_stat=%+v: %v", playerStat, err)
	}

	// set the player stat id
	playerStat.Id = id

	log.Debug("Successfully created player_stat=%+v", playerStat)
	return playerStat, nil
}

///////////////////////////////////////////////////////////////////////////////
// Games

func AllGames() (models.Games, error) {
	var games models.Games
	err := withTx(func(tx *sqlx.Tx) error {
		var tempGames []*game
		log.Debug("Getting all games")
		err := tx.Select(&tempGames, `SELECT * FROM games`)
		if err != nil {
			return fmt.Errorf("Failed to get all games: %v", err)
		}

		for _, game := range tempGames {
			stage, err := GetStage(game.StageId)
			if err != nil {
				return err
			}
			game.Stage = stage

			playerStats, err := getPlayerStatsForGame(tx, game.Id)
			if err != nil {
				return err
			}
			game.PlayerStats = playerStats

			games = append(games, &game.Game)
		}

		log.Debug("Successfully got all games")
		return nil
	})
	return games, err
}

func GetGame(id int) (*models.Game, error) {
	var game game
	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Getting game with id=%d", id)
		err := tx.Get(&game, `SELECT * FROM games WHERE id=$1`, id)
		if err != nil {
			return fmt.Errorf("Failed to get game with id=%q: %v", id, err)
		}

		stage, err := GetStage(game.StageId)
		if err != nil {
			return err
		}
		game.Stage = stage

		playerStats, err := getPlayerStatsForGame(tx, id)
		if err != nil {
			return err
		}
		game.PlayerStats = playerStats

		log.Debug("Successfully got game with id=%d", id)
		return nil
	})
	return &game.Game, err
}

func CreateGame(game *models.Game) (*models.Game, error) {
	stmt := `INSERT INTO games (
		time,
		stage_id,
		player_count
	) VALUES ($1,$2,$3) RETURNING id`

	err := withTx(func(tx *sqlx.Tx) error {
		log.Debug("Creating game=%+v", game)
		var id int
		err := tx.Get(&id, stmt, game.Time.In(utc), game.Stage.Id, game.PlayerCount)
		if err != nil {
			return fmt.Errorf("Failed to create game=%+v: %v", game, err)
		}

		// set game id
		game.Id = id

		for _, playerStat := range game.PlayerStats {
			_, err := createPlayerStat(tx, playerStat, game.Id)
			if err != nil {
				return err
			}
		}

		log.Debug("Successfully created game=%+v", game)
		return nil
	})
	return game, err
}
