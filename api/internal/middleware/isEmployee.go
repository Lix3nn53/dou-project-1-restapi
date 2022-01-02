package middleware

import (
	appError "dou-survey/app/error"
	"dou-survey/app/model/userModel"
	"dou-survey/app/service/employeeService"
	"net/http"

	"github.com/gin-gonic/gin"

	"dou-survey/internal/logger"
)

type isEmployeeMiddleware struct {
	logger logger.Logger
	ec     employeeService.EmployeeServiceInterface
}

//IsEmployeeMiddlewareInterface ...
type IsEmployeeMiddlewareInterface interface {
	Handler() gin.HandlerFunc
}

//NewIsEmployeeMiddleware ...
func NewIsEmployeeMiddleware(logger logger.Logger, ec employeeService.EmployeeServiceInterface) IsEmployeeMiddlewareInterface {
	return &isEmployeeMiddleware{
		logger,
		ec,
	}
}

//Handler ...
func (cm isEmployeeMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// User was added to context in middleware
		var user userModel.User = c.MustGet("user").(userModel.User)

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
