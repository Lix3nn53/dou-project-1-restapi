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
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

//SurveyControllerInterface define the survey controller interface methods
type SurveyControllerInterface interface {
	Vote(c *gin.Context)
	ListActive(c *gin.Context)
	ListResults(c *gin.Context)
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

type VoteRequestBody struct {
	ChoiceID uint `binding:"required"`
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) Vote(c *gin.Context) {
	var requestBody VoteRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	user := c.MustGet("user").(*userModel.UserReduced)
	userFull, err := uc.userService.FindByIdNumber(user.IDNumber)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	vote, err := uc.service.Vote(userFull.ID, requestBody.ChoiceID)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, vote)
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) ListActive(c *gin.Context) {
	limit := c.DefaultQuery("limit", "5")
	offset := c.DefaultQuery("offset", "0")

	limitInt64, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	offsetInt64, err := strconv.ParseUint(offset, 10, 32)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	limitInt := uint(limitInt64)
	offsetInt := uint(offsetInt64)

	result, err := uc.service.List(limitInt, offsetInt)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	// uc.logger.Infof("%#v", result)

	c.JSON(http.StatusOK, result)
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) ListResults(c *gin.Context) {
	limit := c.DefaultQuery("limit", "5")
	offset := c.DefaultQuery("offset", "0")

	limitInt64, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	offsetInt64, err := strconv.ParseUint(offset, 10, 32)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	limitInt := uint(limitInt64)
	offsetInt := uint(offsetInt64)

	result, err := uc.service.ListWithDetails(limitInt, offsetInt)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	// uc.logger.Infof("%#v", result)

	c.JSON(http.StatusOK, result)
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) Info(c *gin.Context) {
	surveyID := c.Param("survey")
	withChoicesQuery := c.DefaultQuery("withChoices", "false") // if survey is active, return withoutVotes instead of Reduced info
	withChoices, err := strconv.ParseBool(withChoicesQuery)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	surveyIDInt64, err := strconv.ParseUint(surveyID, 10, 32)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	surveyIDInt := uint(surveyIDInt64)

	survey, err := uc.service.FindByIDReduced(surveyIDInt)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	now := time.Now().UTC()

	// uc.logger.Info(now)
	// uc.logger.Info(survey.DateStart)
	// uc.logger.Info(survey.DateEnd)
	// uc.logger.Info(now.After(survey.DateStart))
	// uc.logger.Info(now.Before(survey.DateEnd))

	if !(now.After(survey.DateStart) && now.Before(survey.DateEnd)) { // do not enter if survey is active
		// now is before start or now is after end
		// so survey is not active but why?
		if now.After(survey.DateStart) {
			// now is after start which means now is also after end, voting ended, survey completed
			survey, err = uc.service.FindByIDWithVotes(surveyIDInt)
			if err != nil {
				uc.logger.Error(err.Error())
				appError.Respond(c, http.StatusBadRequest, err)
				return
			}
		} // else if now.Before(survey.DateEnd) {
		// 	// now is before end which means now is also before start, survey voting havent started
		// }
	} else if withChoices {
		survey, err = uc.service.FindByIDWithoutVotes(surveyIDInt)
		if err != nil {
			uc.logger.Error(err.Error())
			appError.Respond(c, http.StatusBadRequest, err)
			return
		}
	}

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

	user := c.MustGet("user").(*userModel.UserReduced)
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
