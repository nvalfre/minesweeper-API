package router

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/controllers/game"
	"minesweeper-API/controllers/ping"
)

const pingEndpoint = "/ping"
const newGame = "/game"
const newClickMovement = "/game/movement"

func InitRoutes(r *gin.Engine) {
	r.GET(pingEndpoint, ping.Ping())
	r.POST(newGame, game.Controller.StartNewGame)
	r.POST(newClickMovement, game.Controller.ClickPosition)
}
