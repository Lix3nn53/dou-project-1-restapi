package v1

import (
	"dou-survey/app/controller/surveyController"

	"github.com/gin-gonic/gin"
)

func SetupSurveysRoute(surveys *gin.RouterGroup, c surveyController.SurveyControllerInterface) *gin.RouterGroup {
	surveys.GET("/list", c.List)
	surveys.POST("/create", c.Create)

	return surveys
}
