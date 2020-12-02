package router

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/controllers/game"
	"minesweeper-API/controllers/ping"
)

const pingEndpoint = "/ping"
const newGame = "/game"
const startGame = "/game/start"
const newClickMovement = "/game/movement"

func InitRoutes(r *gin.Engine) {
	r.GET(pingEndpoint, ping.Ping())
	r.POST(newGame, game.Controller.CreateNewGame)
	r.POST(startGame, game.Controller.StartGame)
	r.POST(newClickMovement, game.Controller.ClickPosition)
}
