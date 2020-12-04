package ping

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/utils"
)

func Ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		utils.BuildHeaders(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
