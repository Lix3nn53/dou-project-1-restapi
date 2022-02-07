package controller

import (
	appError "dou-survey/app/error"
	"dou-survey/app/model"
	"dou-survey/app/service"
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
	ChoiceVoters(c *gin.Context)
	Vote(c *gin.Context)
	ListWaitingConfirmation(c *gin.Context)
	ListActive(c *gin.Context)
	ListResults(c *gin.Context)
	CountWaitingConfirmation(c *gin.Context)
	CountActive(c *gin.Context)
	CountResults(c *gin.Context)
	Info(c *gin.Context)
	Create(c *gin.Context)
}

// SurveyController handles communication with the survey service
type SurveyController struct {
	service     service.SurveyServiceInterface
	userService service.UserServiceInterface
	logger      logger.Logger
}

// NewSurveyController implements the survey controller interface.
func NewSurveyController(service service.SurveyServiceInterface, userService service.UserServiceInterface, logger logger.Logger) SurveyControllerInterface {
	return &SurveyController{
		service,
		userService,
		logger,
	}
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) ChoiceVoters(c *gin.Context) {
	choiceID := c.Param("choice")

	choiceIDInt64, err := strconv.ParseUint(choiceID, 10, 32)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	choiceIDInt := uint(choiceIDInt64)

	voters, err := uc.service.ChoiceVotersInfo(choiceIDInt)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, voters)
}

type VoteRequestBody struct {
	SurveyID uint   `binding:"required"`
	Votes    []uint `binding:"required"`
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) Vote(c *gin.Context) {
	var requestBody VoteRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	user := c.MustGet("user").(*model.UserReduced)
	userFull, err := uc.userService.FindByIdNumber(user.IDNumber)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	// check vote count matches
	count, err := uc.service.CountQuestion(requestBody.SurveyID)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}
	if count != len(requestBody.Votes) {
		err := errors.New("vote count does not match")
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	// check voted before
	votedAlready, err := uc.service.VotedAlready(userFull.ID, requestBody.SurveyID)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}
	if votedAlready {
		err := errors.New("vote submitted already for this survey")
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	// submit vote
	created, err := uc.service.Vote(userFull.ID, requestBody.SurveyID, requestBody.Votes)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, created)
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) CountWaitingConfirmation(c *gin.Context) {
	result, err := uc.service.CountWaitingConfirmation()
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) CountActive(c *gin.Context) {
	result, err := uc.service.CountActive()
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) CountResults(c *gin.Context) {
	result, err := uc.service.CountResults()
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Find implements the method to handle the service to find a survey by the primary key
func (uc *SurveyController) ListWaitingConfirmation(c *gin.Context) {
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

	result, err := uc.service.ListWaitingConfirmation(limitInt, offsetInt)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	// uc.logger.Infof("%#v", result)

	c.JSON(http.StatusOK, result)
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

	result, err := uc.service.ListActive(limitInt, offsetInt)
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

	result, err := uc.service.ListResults(limitInt, offsetInt)
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

	if !(now.After(survey.DateStart) && now.Before(survey.DateEnd)) { // if survey is inactive, you cant get votes
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
	} else if withChoices { // if survey is active, do you want choices
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
	var requestBody *model.Survey

	uc.logger.Info(requestBody)

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	user := c.MustGet("user").(*model.UserReduced)
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

	created, err := uc.service.Create(requestBody)

	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, created)
}
