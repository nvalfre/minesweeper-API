package movement

import "github.com/gin-gonic/gin"

var Controller movementControllerInterface = &movementController{}

type movementControllerInterface interface {
	PostMovement(c *gin.Context)
}
type movementController struct {
}

func (controller *movementController) PostMovement(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "newMovement",
	})
}
