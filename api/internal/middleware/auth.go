package middleware

import (
	appError "dou-survey/app/error"
	"dou-survey/app/service"
	"dou-survey/internal/logger"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	logger logger.Logger
	ac     service.AuthServiceInterface
	uc     service.UserServiceInterface
}

//AuthMiddlewareInterface ...
type AuthMiddlewareInterface interface {
	Handler() gin.HandlerFunc
}

//NewAuthMiddleware ...
func NewAuthMiddleware(logger logger.Logger, ac service.AuthServiceInterface, uc service.UserServiceInterface) AuthMiddlewareInterface {
	return &authMiddleware{
		logger,
		ac,
		uc,
	}
}

//Handler ...
func (cm authMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cm.logger.Infof("Auth Middleware: %s", c.ClientIP())

		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			appError.Respond(c, http.StatusBadRequest, errors.New("no authorization header provided"))
			c.Abort()
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		if token == auth {
			appError.Respond(c, http.StatusBadRequest, errors.New("could not find bearer token in authorization header"))
			c.Abort()
			return
		}

		id, err := cm.ac.TokenValidate(token, os.Getenv("ACCESS_SECRET"))
		if err != nil {
			appError.Respond(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		user, err := cm.uc.FindByIDReduced(id)
		if err != nil {
			cm.logger.Error(err.Error())
			appError.Respond(c, http.StatusNotFound, err)
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
