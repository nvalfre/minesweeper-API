package router

import "github.com/gin-gonic/gin"

const pingEndpoint = "/ping"
const userCreateEndpoint = "/user"
const userLogin = "/login"
const userLogout = "/logout"
const newGame = "/game"
const newMovement = "/game/movement"

func InitRoutes() {

	r := gin.Default()

	r.GET(pingEndpoint, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
