package route

import (
	"dou-survey/internal/dic"
	"dou-survey/internal/logger"
	"dou-survey/internal/middleware"
	routev1 "dou-survey/internal/route/v1"
	"dou-survey/internal/storage"
	"fmt"
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
		{
			userService := dic.InitUserService(userRepo)

			// middleware
			authMiddlewareHandler := middleware.NewAuthMiddleware(logger, authService, userService).Handler()
			users.Use(authMiddlewareHandler)

			// route setup
			userCont := dic.InitUserController(userService, logger)
			routev1.SetupUserRoute(users, userCont)

			// create test user
			newUser, err := userService.CreateUser("62236422322", "test@mail.co", "123456")

			if err != nil {
				logger.Error(err.Error())
			}

			logger.Info(newUser)

			// employee, is child of user
			employees := users.Group("/employees")
			{
				employeeRepo := dic.InitEmployeeRepository(db)
				employeeService := dic.InitEmployeeService(employeeRepo)

				// middleware
				isEmployeeMiddleware := middleware.NewIsEmployeeMiddleware(logger, employeeService)
				employees.Use(isEmployeeMiddleware.Handler())

				// route setup
				employeeController := dic.InitEmployeeController(employeeService, logger)
				routev1.SetupEmployeeRoute(employees, employeeController)
			}
		}

	}

	return r
}
