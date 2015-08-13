package models

import (
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// Characters

type Character struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type Characters []*Character

func (c Characters) Len() int           { return len(c) }
func (c Characters) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Characters) Less(i, j int) bool { return c[i].Name < c[j].Name }

///////////////////////////////////////////////////////////////////////////////
// Games

type Game struct {
	Id          int       `db:"id"`
	Time        time.Time `db:"time"`
	PlayerCount int       `db:"player_count"`
	Stage       *Stage
	PlayerStats PlayerStats
}

func (game *Game) Players() Players {
	var players Players
	for _, stat := range game.PlayerStats {
		players = append(players, stat.Player)
	}
	return players
}

func (game *Game) Winner() string {
	var winner string

	for _, playerStats := range game.
	PlayerStats {
		if playerStats.Rank == 1 {
			p := *playerStats.Player
			winner = p.FirstName + " " + p.LastName
			break
		}
	}
	return winner
}

type Games []*Game

func (g Games) Len() int           { return len(g) }
func (g Games) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g Games) Less(i, j int) bool { return g[i].Time.After(g[j].Time) }

///////////////////////////////////////////////////////////////////////////////
// Players

type Player struct {
	Id         int    `db:"id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Email      string `db:"email"`
	Tag        string `db:"tag"`
	IsDisabled bool   `db:"is_disabled"`
}

type Players []*Player

func (p Players) Len() int           { return len(p) }
func (p Players) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Players) Less(i, j int) bool { return p[i].LastName < p[j].LastName }

///////////////////////////////////////////////////////////////////////////////
// Player Stats

type PlayerStat struct {
	Player        *Player
	Character     *Character
	Id            int  `db:"id"`
	IsRandom      bool `db:"is_random"`
	Rank          int  `db:"rank"`
	Number        int  `db:"number"`
	Kills         int  `db:"kills"`
	Deaths        int  `db:"deaths"`
	SelfDestructs int  `db:"self_destructs"`
}

type PlayerStats []*PlayerStat

func (p PlayerStats) Len() int           { return len(p) }
func (p PlayerStats) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PlayerStats) Less(i, j int) bool { return p[i].Number < p[j].Number }

///////////////////////////////////////////////////////////////////////////////
// Stages

type Stage struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type Stages []*Stage

func (s Stages) Len() int           { return len(s) }
func (s Stages) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Stages) Less(i, j int) bool { return s[i].Name < s[j].Name }
