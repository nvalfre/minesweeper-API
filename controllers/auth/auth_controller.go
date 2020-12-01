package auth

import "github.com/gin-gonic/gin"

var Controller authControllerInterface = &authController{}

type authControllerInterface interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}
type authController struct {
}

func (controller *authController) Login(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "newMovement",
	})
}

func (controller *authController) Logout(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "newMovement",
	})
}
