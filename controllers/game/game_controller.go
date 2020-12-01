package game

import "github.com/gin-gonic/gin"

var Controller gameControllerInterface = &gameController{}

type gameControllerInterface interface {
	CreateNewGame(c *gin.Context)
}
type gameController struct {
}

func (controller *gameController) CreateNewGame(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "newMovement",
	})
}
