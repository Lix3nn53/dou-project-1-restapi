package authController

import (
	appError "dou-survey/app/error"
	"dou-survey/app/service/authService"
	"dou-survey/internal/logger"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//UserControllerInterface define the user controller interface methods
type AuthControllerInterface interface {
	Login(c *gin.Context)
	RefreshAccessToken(c *gin.Context)
	Logout(c *gin.Context)
}

// UserController handles communication with the user service
type AuthController struct {
	service authService.AuthServiceInterface
	logger  logger.Logger
}

// NewUserController implements the user controller interface.
func NewAuthController(service authService.AuthServiceInterface, logger logger.Logger) AuthControllerInterface {
	return &AuthController{
		service,
		logger,
	}
}

type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type RefreshAccessResponse struct {
	AccessToken string `json:"access_token"`
}

type LoginRequestBody struct {
	IdentificationNumber string `json:"identificationNumber"`
	Password             string `json:"password"`
}

// Find implements the method to handle the service to find a user by the primary key
func (uc *AuthController) Login(c *gin.Context) {
	var requestBody LoginRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	refreshToken, accessToken, err := uc.service.Login(requestBody.IdentificationNumber, requestBody.Password)
	if err != nil {
		uc.logger.Error(err.Error())
		appError.Respond(c, http.StatusBadRequest, err)
		return
	}

	response := AuthResponse{RefreshToken: refreshToken, AccessToken: accessToken}
	c.JSON(http.StatusOK, response)
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

	response := RefreshAccessResponse{AccessToken: accessToken}
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
