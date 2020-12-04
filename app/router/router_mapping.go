package router

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/controllers/game"
	"minesweeper-API/controllers/ping"
)

const pingEndpoint = "/ping"
const newGame = "/game/new"
const gameHistory = "/game/game_history"
const pauseGame = "/game/pause"
const resumeGame = "/game/resume"
const newClickMovement = "/game/movement"

func InitRoutes(r *gin.Engine) {
	r.GET(pingEndpoint, ping.Ping())
	r.POST(newGame, game.Controller.StartNewGame)
	r.GET(gameHistory, game.Controller.GetHistory)
	r.PUT(newClickMovement, game.Controller.ClickPosition)
	r.PUT(pauseGame, game.Controller.PauseGame)
	r.PUT(resumeGame, game.Controller.ResumeGame)
}
