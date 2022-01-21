package v1

import (
	"dou-survey/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupUserRoute(users *gin.RouterGroup, c controller.UserControllerInterface) *gin.RouterGroup {
	users.GET("/info", c.Info)

	return users
}
