package userService

import (
	"goa-golang/app/model/userModel"
	"goa-golang/app/repository/userRepository"
)

//UserServiceInterface define the user service interface methods
type UserServiceInterface interface {
	FindByIdNumber(id string) (user *userModel.User, err error)
	CreateUser(tckn, email, password string) (user *userModel.User, err error)
}

// billingService handles communication with the user repository
type UserService struct {
	userRepo userRepository.UserRepositoryInterface
}

// NewUserService implements the user service interface.
func NewUserService(userRepo userRepository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepo,
	}
}

func (s *UserService) FindByIdNumber(id string) (user *userModel.User, err error) {
	return s.userRepo.FindByIdNumber(id)
}

func (s *UserService) CreateUser(tckn, email, password string) (user *userModel.User, err error) {
	newUser := userModel.User{
		TCKN:     tckn,
		Email:    email,
		Password: password,
	}

	return s.userRepo.CreateUser(newUser)
}
