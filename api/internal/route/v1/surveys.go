package v1

import (
	"dou-survey/app/controller"

	"github.com/gin-gonic/gin"
)

// Surveys route that does not require auth
func SetupSurveysRoute(surveys *gin.RouterGroup, c controller.SurveyControllerInterface) *gin.RouterGroup {
	surveys.GET("/voter-details/:choice", c.ChoiceVoters)
	surveys.GET("/info/:survey", c.Info)
	surveys.GET("/list/active", c.ListActive)
	surveys.GET("/list/results", c.ListResults)
	surveys.GET("/count/active", c.CountActive)
	surveys.GET("/count/results", c.CountResults)

	return surveys
}
