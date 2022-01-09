package v1

import (
	"dou-survey/app/controller/surveyController"

	"github.com/gin-gonic/gin"
)

func SetupSurveysRoute(survey *gin.RouterGroup, c surveyController.SurveyControllerInterface) *gin.RouterGroup {
	survey.POST("/create", c.Create)

	return survey
}
