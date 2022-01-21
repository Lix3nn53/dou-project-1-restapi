package service

import (
	"dou-survey/app/model"
	"dou-survey/app/repository"
)

//UserServiceInterface define the user service interface methods
type UserServiceInterface interface {
	FindByIDReduced(id uint) (user *model.UserReduced, err error)
	FindByIdNumber(id string) (user *model.User, err error)
	CreateUser(*model.User) (user *model.User, err error)
}

// billingService handles communication with the user repository
type UserService struct {
	userRepo repository.UserRepositoryInterface
}

// NewUserService implements the user service interface.
func NewUserService(userRepo repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepo,
	}
}

func (s *UserService) FindByIDReduced(id uint) (user *model.UserReduced, err error) {
	return s.userRepo.FindByIDReduced(id)
}

func (s *UserService) FindByIdNumber(id string) (user *model.User, err error) {
	return s.userRepo.FindByIdNumber(id)
}

func (s *UserService) CreateUser(newUser *model.User) (user *model.User, err error) {
	return s.userRepo.CreateUser(newUser)
}
