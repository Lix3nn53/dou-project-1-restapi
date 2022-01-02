package middleware

import (
	"github.com/gin-gonic/gin"

	"goa-golang/internal/logger"
)

type isEmployeeMiddleware struct {
	logger logger.Logger
}

//IsEmployeeMiddlewareInterface ...
type IsEmployeeMiddlewareInterface interface {
	Handler() gin.HandlerFunc
}

//NewIsEmployeeMiddleware ...
func NewIsEmployeeMiddleware(logger logger.Logger) IsEmployeeMiddlewareInterface {
	return &isEmployeeMiddleware{
		logger,
	}
}

//Handler ...
func (cm isEmployeeMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cm.logger.Infof("IsEmployee Middleware: %s", c.ClientIP())

		c.Next()
	}
}
