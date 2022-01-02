package v1

import (
	"goa-golang/app/controller/userController"

	"github.com/gin-gonic/gin"
)

func SetupUserRoute(users *gin.RouterGroup, c userController.UserControllerInterface) *gin.RouterGroup {
	users.GET("/info", c.Info)

	return users
}
