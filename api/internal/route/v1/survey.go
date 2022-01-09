package v1

import (
	"dou-survey/app/controller/surveyController"

	"github.com/gin-gonic/gin"
)

func SetupSurveyRoute(survey *gin.RouterGroup, c surveyController.SurveyControllerInterface) *gin.RouterGroup {
	survey.GET("/:survey", c.Info)

	return survey
}
