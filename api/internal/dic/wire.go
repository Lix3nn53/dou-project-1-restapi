//go:build wireinject
// +build wireinject

//
// The build tag makes sure the stub is not built in the final build.
//
//lint:file-ignore U1000 Ignore all unused code

package dic

import (
	"dou-survey/app/controller"
	"dou-survey/app/repository"
	"dou-survey/app/service"
	"dou-survey/internal/logger"
	"dou-survey/internal/storage"

	"github.com/google/wire"
)

// User
func initUserRepository(db *storage.DbStore) repository.UserRepositoryInterface {
	wire.Build(repository.NewUserRepository)

	return &repository.UserRepository{}
}

func initUserService(userRepo repository.UserRepositoryInterface) service.UserServiceInterface {
	wire.Build(service.NewUserService)

	return &service.UserService{}
}

func initUserController(us service.UserServiceInterface, logger logger.Logger) controller.UserControllerInterface {
	wire.Build(controller.NewUserController)

	return &controller.UserController{}
}

// Auth
func initAuthService(userRepo repository.UserRepositoryInterface, logger logger.Logger) service.AuthServiceInterface {
	wire.Build(service.NewAuthService)

	return &service.AuthService{}
}

func initAuthController(as service.AuthServiceInterface, us service.UserServiceInterface, logger logger.Logger) controller.AuthControllerInterface {
	wire.Build(controller.NewAuthController)

	return &controller.AuthController{}
}

// Employee
func initEmployeeRepository(db *storage.DbStore) repository.EmployeeRepositoryInterface {
	wire.Build(repository.NewEmployeeRepository)

	return &repository.EmployeeRepository{}
}

func initEmployeeService(employeeRepo repository.EmployeeRepositoryInterface) service.EmployeeServiceInterface {
	wire.Build(service.NewEmployeeService)

	return &service.EmployeeService{}
}

func initEmployeeController(us service.EmployeeServiceInterface, logger logger.Logger) controller.EmployeeControllerInterface {
	wire.Build(controller.NewEmployeeController)

	return &controller.EmployeeController{}
}

// Survey
func initSurveyRepository(db *storage.DbStore, logger logger.Logger) repository.SurveyRepositoryInterface {
	wire.Build(repository.NewSurveyRepository)

	return &repository.SurveyRepository{}
}

func initSurveyService(surveyRepo repository.SurveyRepositoryInterface) service.SurveyServiceInterface {
	wire.Build(service.NewSurveyService)

	return &service.SurveyService{}
}

func initSurveyController(ss service.SurveyServiceInterface, us service.UserServiceInterface, logger logger.Logger) controller.SurveyControllerInterface {
	wire.Build(controller.NewSurveyController)

	return &controller.SurveyController{}
}

// func initBillingService(db *storage.DbStore) billingService.BillingServiceInterface {
// 	wire.Build(billingRepository.NewBillingRepository, billingService.NewBillingService)

// 	return &billingService.BillingService{}
// }

// func initBillingController(ub billingService.BillingServiceInterface, us userService.UserServiceInterface, logger logger.Logger) billingController.BillingControllerInterface {
// 	wire.Build(billingController.NewBillingController)

// 	return &billingController.BillingController{}
// }
