package controllers

import (
	"net/http"
	"github.com/martini-contrib/render"
)

func Home(render render.Render, request *http.Request) {
	render.HTML(http.StatusOK, "home", "Home")
}

func NotFound(render render.Render, request *http.Request) {
	render.HTML(http.StatusNotFound, "404", "Not Found")
}

func InternalServerError(render render.Render, request *http.Request) {
	render.HTML(http.StatusNotFound, "500", "Server Error")
}
