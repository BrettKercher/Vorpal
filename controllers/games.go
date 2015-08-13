package controllers

import (
	"net/http"
	"fmt"
	"time"
	"sort"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/hudl/vorpal/database"
	"github.com/hudl/vorpal/models"
)

func GetGames(render render.Render, request *http.Request) {
	games, err := database.AllGames()
	if err != nil {
		log.Error(err.Error())
	}

	sort.Sort(games)

	render.HTML(200, "games", games)
}

func GetGameById(render render.Render, request *http.Request, params martini.Params) {
	var id int
	_, err := fmt.Sscanf(params["id"], "%d", &id)
	if err != nil {
		log.Error(err.Error())
	}

	game, err := database.GetGame(id)
	if err != nil {
		log.Error(err.Error())
	}
	render.HTML(http.StatusOK, "game", game)
}

func GetNewGame(render render.Render, request *http.Request) {

	stages, err := database.AllStages()
	if err != nil {
		log.Error(err.Error())
	}
	characters, err := database.AllCharacters()
	if err != nil {
		log.Error(err.Error())
	}
	players, err := database.AllPlayers()
	if err != nil {
		log.Error(err.Error())
	}

	newGame := NewGameModel {
		Stages: 	stages,
		Characters: characters,
		Players: 	players,
	}

	render.HTML(200, "new-game", newGame)
}

func CreateGame(render render.Render, writer http.ResponseWriter, request *http.Request, gameData GameData) {
	///////////////////////////////////////////////////////////////////////////////
	// Get Stage Info
	stage, err := database.GetStageByName(gameData.Stage)
	if err != nil {
		log.Error(err.Error())
	}

	///////////////////////////////////////////////////////////////////////////////
	//Get Player Count
	var playerCount int
	_, err = fmt.Sscanf(gameData.PlayerCount, "%d", &playerCount)
	if err != nil {
		log.Error(err.Error())
	}

	///////////////////////////////////////////////////////////////////////////////
	//Get Player Stats

	playerStats := models.PlayerStats {}

	/**** Player One ****/
	if playerCount >= 1 {
		// Player
		var playerOneId int
		_, err = fmt.Sscanf(gameData.PlayerOneId, "%d", &playerOneId)
		if err != nil {
			log.Error(err.Error())
		}
		player1, err := database.GetPlayer(playerOneId)
		if err != nil {
			log.Error(err.Error())
		}

		// Character
		var characterOneId int
		_, err = fmt.Sscanf(gameData.CharacterOneId, "%d", &characterOneId)
		if err != nil {
			log.Error(err.Error())
		}
		character1, err := database.GetCharacter(characterOneId)
		if err != nil {
			log.Error(err.Error())
		}

		// IsRandom
		var isP1Random bool
		if gameData.IsP1Random == "1" {
			isP1Random = true
		} else {
			isP1Random = false
		}

		// Rank
		var rankOne int
		_, err = fmt.Sscanf(gameData.RankOne, "%d", &rankOne)
		if err != nil {
			log.Error(err.Error())
		}

		// Kills
		var killsOne int
		_, err = fmt.Sscanf(gameData.KillsOne, "%d", &killsOne)
		if err != nil {
			log.Error(err.Error())
		}

		// Deaths
		var deathsOne int
		_, err = fmt.Sscanf(gameData.DeathsOne, "%d", &deathsOne)
		if err != nil {
			log.Error(err.Error())
		}

		// SelfDestructs
		var sdOne int
		_, err = fmt.Sscanf(gameData.SDOne, "%d", &sdOne)
		if err != nil {
			log.Error(err.Error())
		}


		playerStats = append(playerStats,
			&models.PlayerStat {
				Player:    		player1,
				Character: 		character1,
				IsRandom:  		isP1Random,
				Rank:      		rankOne,
				Number:    		1,
				Kills:			killsOne,
				Deaths:			deathsOne,
				SelfDestructs: 	sdOne,
			})
	}

	/**** Player Two ****/
	if playerCount >= 2 {
		// Player
		var playerTwoId int
		_, err = fmt.Sscanf(gameData.PlayerTwoId, "%d", &playerTwoId)
		if err != nil {
			log.Error(err.Error())
		}
		player2, err := database.GetPlayer(playerTwoId)
		if err != nil {
			log.Error(err.Error())
		}

		// Character
		var characterTwoId int
		_, err = fmt.Sscanf(gameData.CharacterTwoId, "%d", &characterTwoId)
		if err != nil {
			log.Error(err.Error())
		}
		character2, err := database.GetCharacter(characterTwoId)
		if err != nil {
			log.Error(err.Error())
		}

		// IsRandom
		var isP2Random bool
		if gameData.IsP2Random == "1" {
			isP2Random = true
		} else {
			isP2Random = false
		}

		// Rank
		var rankTwo int
		_, err = fmt.Sscanf(gameData.RankTwo, "%d", &rankTwo)
		if err != nil {
			log.Error(err.Error())
		}

		// Kills
		var killsTwo int
		_, err = fmt.Sscanf(gameData.KillsTwo, "%d", &killsTwo)
		if err != nil {
			log.Error(err.Error())
		}

		// Deaths
		var deathsTwo int
		_, err = fmt.Sscanf(gameData.DeathsTwo, "%d", &deathsTwo)
		if err != nil {
			log.Error(err.Error())
		}

		// SelfDestructs
		var sdTwo int
		_, err = fmt.Sscanf(gameData.SDTwo, "%d", &sdTwo)
		if err != nil {
			log.Error(err.Error())
		}

		playerStats = append(playerStats,
			&models.PlayerStat {
				Player:    		player2,
				Character: 		character2,
				IsRandom:  		isP2Random,
				Rank:      		rankTwo,
				Number:    		2,
				Kills:			killsTwo,
				Deaths:			deathsTwo,
				SelfDestructs: 	sdTwo,
			})
	}

	/**** Player Three ****/
	if playerCount >=3 {
		// Player
		var playerThreeId int
		_, err = fmt.Sscanf(gameData.PlayerThreeId, "%d", &playerThreeId)
		if err != nil {
			log.Error(err.Error())
		}
		player3, err := database.GetPlayer(playerThreeId)
		if err != nil {
			log.Error(err.Error())
		}

		// Character
		var characterThreeId int
		_, err = fmt.Sscanf(gameData.CharacterThreeId, "%d", &characterThreeId)
		if err != nil {
			log.Error(err.Error())
		}
		character3, err := database.GetCharacter(characterThreeId)
		if err != nil {
			log.Error(err.Error())
		}

		// IsRandom
		var isP3Random bool
		if gameData.IsP3Random == "1" {
			isP3Random = true
		} else {
			isP3Random = false
		}

		// Rank
		var rankThree int
		_, err = fmt.Sscanf(gameData.RankThree, "%d", &rankThree)
		if err != nil {
			log.Error(err.Error())
		}

		// Kills
		var killsThree int
		_, err = fmt.Sscanf(gameData.KillsThree, "%d", &killsThree)
		if err != nil {
			log.Error(err.Error())
		}

		// Deaths
		var deathsThree int
		_, err = fmt.Sscanf(gameData.DeathsThree, "%d", &deathsThree)
		if err != nil {
			log.Error(err.Error())
		}

		// SelfDestructs
		var sdThree int
		_, err = fmt.Sscanf(gameData.SDThree, "%d", &sdThree)
		if err != nil {
			log.Error(err.Error())
		}

		playerStats = append(playerStats,
			&models.PlayerStat {
				Player:    		player3,
				Character: 		character3,
				IsRandom:  		isP3Random,
				Rank:      		rankThree,
				Number:    		3,
				Kills:			killsThree,
				Deaths:			deathsThree,
				SelfDestructs: 	sdThree,
			})
	}

	/**** Player Four ****/

	if playerCount == 4 {
		// Player
		var playerFourId int
		_, err = fmt.Sscanf(gameData.PlayerFourId, "%d", &playerFourId)
		if err != nil {
			log.Error(err.Error())
		}
		player4, err := database.GetPlayer(playerFourId)
		if err != nil {
			log.Error(err.Error())
		}

		// Character
		var characterFourId int
		_, err = fmt.Sscanf(gameData.CharacterFourId, "%d", &characterFourId)
		if err != nil {
			log.Error(err.Error())
		}
		character4, err := database.GetCharacter(characterFourId)
		if err != nil {
			log.Error(err.Error())
		}

		// IsRandom
		var isP4Random bool
		if gameData.IsP4Random == "1" {
			isP4Random = true
		} else {
			isP4Random = false
		}

		// Rank
		var rankFour int
		_, err = fmt.Sscanf(gameData.RankFour, "%d", &rankFour)
		if err != nil {
			log.Error(err.Error())
		}

		// Kills
		var killsFour int
		_, err = fmt.Sscanf(gameData.KillsFour, "%d", &killsFour)
		if err != nil {
			log.Error(err.Error())
		}

		// Deaths
		var deathsFour int
		_, err = fmt.Sscanf(gameData.DeathsFour, "%d", &deathsFour)
		if err != nil {
			log.Error(err.Error())
		}

		// SelfDestructs
		var sdFour int
		_, err = fmt.Sscanf(gameData.SDFour, "%d", &sdFour)
		if err != nil {
			log.Error(err.Error())
		}

		playerStats = append(playerStats,
			&models.PlayerStat {
				Player:    		player4,
				Character: 		character4,
				IsRandom:  		isP4Random,
				Rank:      		rankFour,
				Number:    		4,
				Kills:			killsFour,
				Deaths:			deathsFour,
				SelfDestructs: 	sdFour,
			})
	}

	gameModel := models.Game {
		Stage: stage,
		PlayerCount: playerCount,
		Time: time.Now(),
		PlayerStats: playerStats,
	}

	_, err = database.CreateGame(&gameModel)
	if err != nil {
		log.Error(err.Error())
	}

	render.Redirect("/games")
}

type GameData struct {
	Stage 			string`form:"stage"`
	PlayerCount		string`form:"playerCount"`
	// P1
	PlayerOneId 	string`form:"playerOneId"`
	CharacterOneId 	string`form:"characterOneId"`
	IsP1Random		string`form:"isP1Random"`
	RankOne			string`form:"optradio1"`
	KillsOne		string`form:"p1Kills"`
	DeathsOne		string`form:"p1Deaths"`
	SDOne			string`form:"p1SDs"`
	// P2
	PlayerTwoId 	string`form:"playerTwoId"`
	CharacterTwoId 	string`form:"characterTwoId"`
	IsP2Random		string`form:"isP2Random"`
	RankTwo 		string`form:"optradio2"`
	KillsTwo		string`form:"p2Kills"`
	DeathsTwo		string`form:"p2Deaths"`
	SDTwo			string`form:"p2SDs"`
	// P3
	PlayerThreeId 		string`form:"playerThreeId"`
	CharacterThreeId 	string`form:"characterThreeId"`
	IsP3Random			string`form:"isP3Random"`
	RankThree 			string`form:"optradio3"`
	KillsThree			string`form:"p3Kills"`
	DeathsThree			string`form:"p3Deaths"`
	SDThree				string`form:"p3SDs"`
	// P4
	PlayerFourId 		string`form:"playerFourId"`
	CharacterFourId 	string`form:"characterFourId"`
	IsP4Random			string`form:"isP4Random"`
	RankFour 			string`form:"optradio4"`
	KillsFour			string`form:"p4Kills"`
	DeathsFour			string`form:"p4Deaths"`
	SDFour				string`form:"p4SDs"`
}

type NewGameModel struct {
	Stages		models.Stages
	Characters 	models.Characters
	Players		models.Players
}