package v1

import (
	"dou-survey/app/controller"

	"github.com/gin-gonic/gin"
)

// Surveys route that does not require auth
func SetupAdminSurveyRoute(surveys *gin.RouterGroup, c controller.SurveyControllerInterface) *gin.RouterGroup {
	surveys.GET("/confirm/:survey", c.GetConfirmed)
	surveys.POST("/confirm", c.Confirm)

	return surveys
}
