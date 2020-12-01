package router

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/controllers/auth"
	"minesweeper-API/controllers/game"
	"minesweeper-API/controllers/movement"
	"minesweeper-API/controllers/ping"
	"minesweeper-API/controllers/user"
)

const pingEndpoint = "/ping"
const userCreateEndpoint = "/user"
const userLogin = "/login"
const userLogout = "/logout"
const newGame = "/game"
const newMovement = "/game/movement"

func InitRoutes(r *gin.Engine) {

	r.GET(pingEndpoint, ping.Ping())
	r.POST(userCreateEndpoint, user.Controller.CreateUser)
	r.POST(userLogin, auth.Controller.Login)
	r.POST(userLogout, auth.Controller.Logout)
	r.POST(newGame, game.Controller.CreateNewGame)
	r.POST(newMovement, movement.Controller.PostMovement)

}
