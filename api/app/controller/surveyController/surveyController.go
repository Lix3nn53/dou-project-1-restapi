package surveyController

import (
	appError "dou-survey/app/error"
	"dou-survey/app/service/surveyService"
	"dou-survey/internal/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SurveyControllerInterface define the survey controller interface methods
type SurveyControllerInterface interface {
	Info(c *gin.Context)
}

// SurveyController handles communication with the survey service
type SurveyController struct {
	service surveyService.SurveyServiceInterface
	logger  logger.Logger
}

// NewSurveyController implements the survey controller interface.
func NewSurveyController(service surveyService.SurveyServiceInterface, logger logger.Logger) SurveyControllerInterface {
	return &SurveyController{
		service,
		logger,
	}
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) Info(c *gin.Context) {
	// Survey was added to context in middleware
	survey := c.Param("survey")

	u64, err := strconv.ParseUint(survey, 10, 32)

	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	uc.service.FindByID(uint(u64))

	c.JSON(http.StatusOK, survey)
}
