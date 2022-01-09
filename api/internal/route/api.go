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
		userService := dic.InitUserService(userRepo)
		authCont := dic.InitAuthController(authService, userService, logger)

		auth := v1.Group("/auth")
		{

			routev1.SetupAuthRoute(auth, authCont)
		}

		// users
		users := v1.Group("/users")
		// middleware
		authMiddlewareHandler := middleware.NewAuthMiddleware(logger, authService, userService).Handler()
		users.Use(authMiddlewareHandler)
		{
			// route setup
			userCont := dic.InitUserController(userService, logger)
			routev1.SetupUserRoute(users, userCont)
		}

		// employees
		employees := v1.Group("/employees")
		{
			// auth middleware
			employees.Use(authMiddlewareHandler)

			employeeRepo := dic.InitEmployeeRepository(db)
			employeeService := dic.InitEmployeeService(employeeRepo)

			// isEmployee middleware
			isEmployeeMiddlewareHandler := middleware.NewIsEmployeeMiddleware(logger, employeeService).Handler()
			employees.Use(isEmployeeMiddlewareHandler)

			// route setup
			employeeController := dic.InitEmployeeController(employeeService, logger)
			routev1.SetupEmployeeRoute(employees, employeeController)
		}

		// survey
		survey := v1.Group("/survey")
		{
			// auth middleware
			survey.Use(authMiddlewareHandler)

			surveyRepo := dic.InitSurveyRepository(db)
			surveyService := dic.InitSurveyService(surveyRepo)

			// route setup
			surveyController := dic.InitSurveyController(surveyService, logger)
			routev1.SetupSurveyRoute(survey, surveyController)
		}
	}

	return r
}
