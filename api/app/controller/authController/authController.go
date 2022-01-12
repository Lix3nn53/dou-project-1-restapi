package authController

import (
	appError "dou-survey/app/error"
	"dou-survey/app/model/userModel"
	"dou-survey/app/service/authService"
	"dou-survey/app/service/userService"
	"dou-survey/internal/logger"
	"errors"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

//UserControllerInterface define the user controller interface methods
type AuthControllerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshAccessToken(c *gin.Context)
	Logout(c *gin.Context)
}

// UserController handles communication with the user service
type AuthController struct {
	service     authService.AuthServiceInterface
	logger      logger.Logger
	userService userService.UserServiceInterface
}

// NewUserController implements the user controller interface.
func NewAuthController(service authService.AuthServiceInterface, userService userService.UserServiceInterface, logger logger.Logger) AuthControllerInterface {
	return &AuthController{
		service,
		logger,
		userService,
	}
}

// Find implements the method to handle the service to find a user by the primary key
func (uc *AuthController) Register(c *gin.Context) {
	var requestBody userModel.User

	uc.logger.Info(requestBody)

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

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

	newUser, err := uc.userService.CreateUser(&requestBody)

	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	uc.logger.Info(newUser)

	c.Status(http.StatusCreated)
}

type LoginRequestBody struct {
	IDNumber string `binding:"required"`
	Password string `binding:"required"`
}

type LoginResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// Find implements the method to handle the service to find a user by the primary key
func (uc *AuthController) Login(c *gin.Context) {
	var requestBody LoginRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	refreshToken, accessToken, err := uc.service.Login(requestBody.IDNumber, requestBody.Password)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	response := LoginResponse{RefreshToken: refreshToken, AccessToken: accessToken}
	c.JSON(http.StatusOK, response)
}

type RefreshAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// Find implements the method to handle the service to find a user by the primary key
func (uc *AuthController) RefreshAccessToken(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	if auth == "" {
		appError.Respond(c, http.StatusBadRequest, errors.New("no authorization header provided"))
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		appError.Respond(c, http.StatusBadRequest, errors.New("could not find bearer token in authorization header"))
		return
	}

	id, err := uc.service.TokenValidateRefresh(token)
	if err != nil {
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	accessToken, err := uc.service.TokenBuildAccess(id)
	if err != nil {
		appError.Respond(c, http.StatusInternalServerError, err)
		return
	}

	response := RefreshAccessTokenResponse{AccessToken: accessToken}
	c.JSON(http.StatusOK, response)
}

// Find implements the method to handle the service to find a user by the primary key
func (uc *AuthController) Logout(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")

	if auth == "" {
		appError.Respond(c, http.StatusBadRequest, errors.New("no authorization header provided"))
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		appError.Respond(c, http.StatusBadRequest, errors.New("could not find bearer token in authorization header"))
		return
	}

	id, err := uc.service.TokenValidateRefresh(token)
	if err != nil {
		appError.Respond(c, http.StatusUnauthorized, err)
		return
	}

	err = uc.service.Logout(id, token)
	if err != nil {
		appError.Respond(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
