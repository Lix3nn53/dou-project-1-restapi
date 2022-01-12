package v1

import (
	"dou-survey/app/controller/surveyController"

	"github.com/gin-gonic/gin"
)

// Surveys route that does not require auth
func SetupSurveysRoute(surveys *gin.RouterGroup, c surveyController.SurveyControllerInterface) *gin.RouterGroup {
	surveys.GET("/info/:survey", c.Info)
	surveys.GET("/list", c.List)

	return surveys
}
