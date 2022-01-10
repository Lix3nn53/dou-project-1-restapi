package surveyController

import (
	appError "dou-survey/app/error"
	"dou-survey/app/model/surveyModel"
	"dou-survey/app/model/userModel"
	"dou-survey/app/service/surveyService"
	"dou-survey/app/service/userService"
	"dou-survey/internal/logger"
	"errors"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

//SurveyControllerInterface define the survey controller interface methods
type SurveyControllerInterface interface {
	List(c *gin.Context)
	Info(c *gin.Context)
	Create(c *gin.Context)
}

// SurveyController handles communication with the survey service
type SurveyController struct {
	service     surveyService.SurveyServiceInterface
	userService userService.UserServiceInterface
	logger      logger.Logger
}

// NewSurveyController implements the survey controller interface.
func NewSurveyController(service surveyService.SurveyServiceInterface, userService userService.UserServiceInterface, logger logger.Logger) SurveyControllerInterface {
	return &SurveyController{
		service,
		userService,
		logger,
	}
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) List(c *gin.Context) {
	limit := c.DefaultQuery("limit", "5")
	offset := c.DefaultQuery("offset", "0")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	result, err := uc.service.List(limitInt, offsetInt)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	uc.logger.Infof("%#v", result)

	c.JSON(http.StatusOK, result)
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) Info(c *gin.Context) {
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

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) Create(c *gin.Context) {
	var requestBody *surveyModel.Survey

	uc.logger.Info(requestBody)

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	user := c.MustGet("user").(*userModel.User)
	userFull, err := uc.userService.FindByIdNumber(user.IDNumber)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}
	requestBody.UserRefer = userFull.ID

	valid, err := govalidator.ValidateStruct(requestBody)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}
	if !valid {
		err := errors.New("fields are not valid")
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	uc.logger.Info("SURVEY CREATE")
	uc.logger.Infof("%#v", requestBody)

	_, err = uc.service.Create(requestBody)

	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
}
