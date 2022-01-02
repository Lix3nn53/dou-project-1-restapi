package route

import (
	"fmt"
	"goa-golang/internal/dic"
	"goa-golang/internal/logger"
	"goa-golang/internal/middleware"
	routev1 "goa-golang/internal/route/v1"
	"goa-golang/internal/storage"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Setup returns initialized routes.
func Setup(db *storage.DbStore, dbCache *storage.DbCache, logger logger.Logger) *gin.Engine {
	// ac := container.Get(dic.AuthController).(controller.AuthControllerInterface)

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())

	// Middleware initialization
	corsMiddleware := middleware.NewCorsMiddleware()
	r.Use(corsMiddleware.Handler())

	// server Routes
	SetupServerRoute(r)

	// v1 Routes
	v1 := r.Group("/v1")
	{
		routev1.SetupDocsRoute(v1)

		// user
		userRepo := dic.InitUserRepository(db)

		// auth
		authService := dic.InitAuthService(userRepo, logger)
		authCont := dic.InitAuthController(authService, logger)

		auth := v1.Group("/auth")
		{

			routev1.SetupAuthRoute(auth, authCont)
		}

		// users
		users := v1.Group("/users")
		userService := dic.InitUserService(userRepo)
		authMiddlewareHandler := middleware.NewAuthMiddleware(logger).Handler(authService, userService) //also used in admin
		{
			userCont := dic.InitUserController(userService, logger)

			// middleware
			users.Use(authMiddlewareHandler)

			routev1.SetupUserRoute(users, userCont)

			newUser, err := userService.CreateUser("62236422322", "test@mail.co", "123456")

			if err != nil {
				logger.Error(err.Error())
			}

			logger.Info(newUser)
		}

		// admin
		admin := v1.Group("/admin")
		{
			// middleware
			admin.Use(authMiddlewareHandler)
			isEmployeeMiddleware := middleware.NewIsEmployeeMiddleware(logger)
			admin.Use(isEmployeeMiddleware.Handler())

			// routev1.SetupUserRoute(users, userCont, authCont)

			// newUser, err := userService.CreateUser("62236422322", "test@mail.co", "123456")

			// if err != nil {
			// 	logger.Error(err.Error())
			// }

			// logger.Info(newUser)
		}
	}

	return r
}
