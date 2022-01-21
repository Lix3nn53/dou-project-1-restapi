package v1

import (
	"dou-survey/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(auth *gin.RouterGroup, c controller.AuthControllerInterface) *gin.RouterGroup {
	auth.GET("refresh_token", c.RefreshAccessToken)
	auth.GET("logout", c.Logout)
	auth.POST("login", c.Login)
	auth.POST("register", c.Register)

	return auth
}
