package middleware

import (
	"errors"
	appError "goa-golang/app/error"
	"goa-golang/app/service/authService"
	"goa-golang/app/service/userService"
	"goa-golang/internal/logger"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	logger logger.Logger
}

//AuthMiddlewareInterface ...
type AuthMiddlewareInterface interface {
	Handler(authService.AuthServiceInterface, userService.UserServiceInterface) gin.HandlerFunc
}

//NewAuthMiddleware ...
func NewAuthMiddleware(logger logger.Logger) AuthMiddlewareInterface {
	return &authMiddleware{
		logger,
	}
}

//Handler ...
func (cm authMiddleware) Handler(ac authService.AuthServiceInterface, uc userService.UserServiceInterface) gin.HandlerFunc {
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

		id, err := ac.TokenValidate(token, os.Getenv("ACCESS_SECRET"))
		if err != nil {
			appError.Respond(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		user, err := uc.FindByIDReduced(id)
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
