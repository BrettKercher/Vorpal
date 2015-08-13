package controllers

import (
	"net/http"
	"fmt"


	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/hudl/vorpal/models"
	"github.com/hudl/vorpal/database"
)

func GetPlayers(render render.Render, request *http.Request) {
	players, err := database.AllPlayers()
	if err != nil {
		log.Error(err.Error())
	}

	render.HTML(http.StatusOK, "players", players)
}

func GetPlayerById(render render.Render, request *http.Request, params martini.Params) {
	var id int
	_, err := fmt.Sscanf(params["id"], "%d", &id)
	if err != nil {
		log.Error(err.Error())
	}

	player, err := database.GetPlayer(id)
	if err != nil {
		log.Error(err.Error())
	}

	stats, err := database.GetPlayerAllTimeStats(id)
	if err != nil {
		log.Error(err.Error())
	}

	// rank, err := database.GetPlayerRank(id)
	// if err != nil {
	// 	 log.Error(err.Error())
	// }

	playerProfile := PlayerProfile {
		Player: *player,
		Stats: *stats,
		//Rank: *rank,
	}

	render.HTML(http.StatusOK, "player", playerProfile)
}

func GetNewPlayer(render render.Render, request *http.Request) {
	render.HTML(200, "new-player", "")
}

func CreatePlayer(render render.Render, writer http.ResponseWriter, request *http.Request, playerData PlayerData) {
	playerModel := models.Player {
		FirstName: playerData.FirstName,
		LastName: playerData.LastName,
		Tag: playerData.Tag,
		Email: playerData.Email,
	}

	_, err := database.CreatePlayer(&playerModel)
	if err != nil {
		log.Error(err.Error())
	}

	render.Redirect("/players")
}

func ModifyPlayer(render render.Render, writer http.ResponseWriter, params martini.Params, request *http.Request, playerData PlayerData, originalData OriginalData) {
	var id int
	_, err := fmt.Sscanf(params["id"], "%d", &id)
	if err != nil {
		log.Error(err.Error())
	}

	if playerData.FirstName != originalData.FirstName || playerData.LastName != originalData.LastName {
		err := database.UpdatePlayerName(id, playerData.FirstName, playerData.LastName)
		if err != nil {
			log.Error(err.Error())
		}
	}

	if playerData.Tag != originalData.Tag {
		err := database.UpdatePlayerTag(id, playerData.Tag)
		if err != nil {
			log.Error(err.Error())
		}
	}

	render.Redirect("/player/" + params["id"])
}

type PlayerProfile struct {
	Player 		models.Player
	Stats 		database.PlayerAllTimeStats
	// Rank 		models.PlayerRank
}

type PlayerData struct {
	FirstName 	string `form:"firstName"`
	LastName 	string `form:"lastName"`
	Tag 		string `form:"tag"`
	Email 		string `form:"email"`
}

type OriginalData struct {
	FirstName 	string `form:"firstNameOriginal"`
	LastName 	string `form:"lastNameOriginal"`
	Tag 		string `form:"tagOriginal"`
}