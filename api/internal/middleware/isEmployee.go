package middleware

import (
	appError "dou-survey/app/error"
	"dou-survey/app/model"
	"dou-survey/app/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"dou-survey/internal/logger"
)

type isEmployeeMiddleware struct {
	logger logger.Logger
	ec     service.EmployeeServiceInterface
	us     service.UserServiceInterface
}

//IsEmployeeMiddlewareInterface ...
type IsEmployeeMiddlewareInterface interface {
	Handler() gin.HandlerFunc
}

//NewIsEmployeeMiddleware ...
func NewIsEmployeeMiddleware(logger logger.Logger, ec service.EmployeeServiceInterface, us service.UserServiceInterface) IsEmployeeMiddlewareInterface {
	return &isEmployeeMiddleware{
		logger,
		ec,
		us,
	}
}

//Handler ...
func (cm isEmployeeMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// User was added to context in middleware
		var userReduced *model.UserReduced = c.MustGet("user").(*model.UserReduced)

		user, err := cm.us.FindByIdNumber(userReduced.IDNumber)
		if err != nil {
			cm.logger.Error(err.Error())
			appError.Respond(c, http.StatusNotFound, err)
			c.Abort()
			return
		}
		employee, err := cm.ec.FindByUserId(user.ID)
		if err != nil {
			cm.logger.Error(err.Error())
			appError.Respond(c, http.StatusNotFound, err)
			c.Abort()
			return
		}

		c.Set("employee", employee)

		c.Next()
	}
}
