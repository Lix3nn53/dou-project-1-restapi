package v1

import (
	"dou-survey/app/controller/surveyController"

	"github.com/gin-gonic/gin"
)

// Survey route that requires auth
func SetupSurveyRoute(survey *gin.RouterGroup, c surveyController.SurveyControllerInterface) *gin.RouterGroup {
	survey.POST("/create", c.Create)
	survey.POST("/vote", c.Vote)

	return survey
}
