package v1

import (
	"dou-survey/app/controller"

	"github.com/gin-gonic/gin"
)

// Surveys route that does not require auth
func SetupAdminSurveysRoute(surveys *gin.RouterGroup, c controller.SurveyControllerInterface) *gin.RouterGroup {
	surveys.GET("/count/waiting", c.CountWaitingConfirmation)
	surveys.GET("/list/waiting", c.ListWaitingConfirmation)

	return surveys
}
