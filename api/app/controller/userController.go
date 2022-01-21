package controller

import (
	"dou-survey/app/service"
	"dou-survey/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UserControllerInterface define the user controller interface methods
type UserControllerInterface interface {
	Info(c *gin.Context)
}

// UserController handles communication with the user service
type UserController struct {
	service service.UserServiceInterface
	logger  logger.Logger
}

// NewUserController implements the user controller interface.
func NewUserController(service service.UserServiceInterface, logger logger.Logger) UserControllerInterface {
	return &UserController{
		service,
		logger,
	}
}

// Find implements the method to handle the service to find a user by the primary key
func (uc *UserController) Info(c *gin.Context) {
	// User was added to context in middleware
	user := c.MustGet("user")

	c.JSON(http.StatusOK, user)
}
