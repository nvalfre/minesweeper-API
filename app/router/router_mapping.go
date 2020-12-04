package router

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/controllers/game"
	"minesweeper-API/controllers/ping"
)

const pingEndpoint = "/ping"
const newGame = "/game"
const pauseGame = "/game/pause"
const resumeGame = "/game/resume"
const newClickMovement = "/game/movement"

func InitRoutes(r *gin.Engine) {
	r.GET(pingEndpoint, ping.Ping())
	r.POST(newGame, game.Controller.StartNewGame)
	r.PUT(newClickMovement, game.Controller.ClickPosition)
	r.PUT(pauseGame, game.Controller.PauseGame)
	r.PUT(resumeGame, game.Controller.ResumeGame)
}
