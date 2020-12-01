package user

import "github.com/gin-gonic/gin"

var Controller userControllerInterface = &userController{}

type userControllerInterface interface {
	CreateUser(c *gin.Context)
}
type userController struct {
}

func (controller *userController) CreateUser(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "newMovement",
	})
}
