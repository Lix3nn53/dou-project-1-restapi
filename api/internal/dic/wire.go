//go:build wireinject
// +build wireinject

//
// The build tag makes sure the stub is not built in the final build.
//
//lint:file-ignore U1000 Ignore all unused code

package dic

import (
	"dou-survey/app/controller/authController"
	"dou-survey/app/controller/employeeController"
	"dou-survey/app/controller/userController"
	"dou-survey/app/repository/employeeRepository"
	"dou-survey/app/repository/userRepository"
	"dou-survey/app/service/authService"
	"dou-survey/app/service/employeeService"
	"dou-survey/app/service/userService"
	"dou-survey/internal/logger"
	"dou-survey/internal/storage"

	"github.com/google/wire"
)

// User
func initUserRepository(db *storage.DbStore) userRepository.UserRepositoryInterface {
	wire.Build(userRepository.NewUserRepository)

	return &userRepository.UserRepository{}
}

func initUserService(userRepo userRepository.UserRepositoryInterface) userService.UserServiceInterface {
	wire.Build(userService.NewUserService)

	return &userService.UserService{}
}

func initUserController(us userService.UserServiceInterface, logger logger.Logger) userController.UserControllerInterface {
	wire.Build(userController.NewUserController)

	return &userController.UserController{}
}

// Auth
func initAuthService(userRepo userRepository.UserRepositoryInterface, logger logger.Logger) authService.AuthServiceInterface {
	wire.Build(authService.NewAuthService)

	return &authService.AuthService{}
}

func initAuthController(us authService.AuthServiceInterface, logger logger.Logger) authController.AuthControllerInterface {
	wire.Build(authController.NewAuthController)

	return &authController.AuthController{}
}

// Employee
func initEmployeeRepository(db *storage.DbStore) employeeRepository.EmployeeRepositoryInterface {
	wire.Build(employeeRepository.NewEmployeeRepository)

	return &employeeRepository.EmployeeRepository{}
}

func initEmployeeService(employeeRepo employeeRepository.EmployeeRepositoryInterface) employeeService.EmployeeServiceInterface {
	wire.Build(employeeService.NewEmployeeService)

	return &employeeService.EmployeeService{}
}

func initEmployeeController(us employeeService.EmployeeServiceInterface, logger logger.Logger) employeeController.EmployeeControllerInterface {
	wire.Build(employeeController.NewEmployeeController)

	return &employeeController.EmployeeController{}
}

// func initBillingService(db *storage.DbStore) billingService.BillingServiceInterface {
// 	wire.Build(billingRepository.NewBillingRepository, billingService.NewBillingService)

// 	return &billingService.BillingService{}
// }

// func initBillingController(ub billingService.BillingServiceInterface, us userService.UserServiceInterface, logger logger.Logger) billingController.BillingControllerInterface {
// 	wire.Build(billingController.NewBillingController)

// 	return &billingController.BillingController{}
// }
