package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

	golog "github.com/op/go-logging"
	stdlog "log"

	"github.com/hudl/vorpal/config"
	"github.com/hudl/vorpal/controllers"
	"github.com/hudl/vorpal/database"
	"github.com/hudl/vorpal/middleware"
)

const logPath = "/var/log/vorpal"

var log = golog.MustGetLogger("main")

func main() {
	configPath := flag.String("c",
		"/etc/vorpal.conf",
		"The path to the config file, defaults to /etc/vorpal.conf")
	flag.Parse()

	log.Info("Started with config=%q", *configPath)
	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Panic(err)
		panic(err.Error())
	}

	log.Info("Configuring logging")
	configureLogging(config.Server.ColorLogs)

	// open new database connection
	database.Open(config.Database.Url())
	defer database.Close()

	log.Info("Creating new router")
	router := martini.NewRouter()

	log.Info("Configuring router")
	configureRoutes(router)

	log.Info("Creating new martini server")
	server := martini.New()

	log.Info("Configuring martini server")
	configureServer(server, router)

	log.Info("Starting vorpal on port=%d", config.Server.Port)
	log.Notice("\"The vorpal blade went snicker-snack!\"")
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), server)
	if err != nil {
		log.Panic(err)
		panic(err.Error())
	}
}

func configureLogging(colorLogs bool) {
	golog.SetFormatter(golog.MustStringFormatter("[0x%{id:x}] [%{level}] [%{module}] %{message}"))
	stdoutLogBackend := golog.NewLogBackend(os.Stdout, "", stdlog.LstdFlags|stdlog.Lshortfile)
	stdoutLogBackend.Color = colorLogs

	golog.SetBackend(stdoutLogBackend)
	golog.SetLevel(golog.DEBUG, "")

	// test logs
	jabberwocky, err := ioutil.ReadFile("the_jabberwocky.txt")
	if err != nil {
		log.Warning("Failed to read \"The Jabberwocky\", it looks like nonsense")
		return
	}
	log.Notice(fmt.Sprintf("\n\n%s", string(jabberwocky)))
}

func configureServer(server *martini.Martini, router martini.Router) {
	server.Use(martini.Recovery())
	server.Use(middleware.Logger())
	server.Use(martini.Static("templates/public", martini.StaticOptions{SkipLogging: true}))
	server.Use(martini.Static("templates/images", martini.StaticOptions{Prefix: "images", SkipLogging: true}))
	server.Use(martini.Static("templates/styles", martini.StaticOptions{Prefix: "styles", SkipLogging: true}))
	server.Use(martini.Static("templates/scripts", martini.StaticOptions{Prefix: "scripts", SkipLogging: true}))
	server.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	server.MapTo(router, (*martini.Routes)(nil))
	server.Action(router.Handle)
}

func configureRoutes(router martini.Router) {
	// 404 handler
	router.NotFound(func(render render.Render) {
		render.Redirect("/404")
	})

	// Common Routes
	router.Get("/", controllers.Home)
	router.Get("/404", controllers.NotFound)
	router.Get("/500", controllers.InternalServerError)

	// Player Routes
	router.Get("/players", controllers.GetPlayers)
	router.Get("/player/new", controllers.GetNewPlayer)
	router.Post("/player/new", binding.Form(controllers.PlayerData{}), controllers.CreatePlayer)
	router.Get("/player/:id", controllers.GetPlayerById)
	router.Post("/player/:id", binding.Form(controllers.PlayerData{}), binding.Form(controllers.OriginalData{}), controllers.ModifyPlayer)

	// Game Routes
	router.Get("/games", controllers.GetGames)
	router.Get("/game/new", controllers.GetNewGame)
	router.Post("/game/new", binding.Form(controllers.GameData{}), controllers.CreateGame)
	router.Get("/game/:id", controllers.GetGameById)
}
