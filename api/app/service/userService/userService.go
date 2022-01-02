package userService

import (
	"dou-survey/app/model/userModel"
	"dou-survey/app/repository/userRepository"

	"gorm.io/datatypes"
)

//UserServiceInterface define the user service interface methods
type UserServiceInterface interface {
	FindByIDReduced(id uint) (user *userModel.User, err error)
	FindByIdNumber(id string) (user *userModel.User, err error)
	CreateUser(id_number, email, password, nationality string,
		birthSex userModel.BirthSex, genderIdentity userModel.GenderIdentity,
		birthDate datatypes.Date) (user *userModel.User, err error)
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

func (s *UserService) FindByIDReduced(id uint) (user *userModel.User, err error) {
	return s.userRepo.FindByIDReduced(id)
}

func (s *UserService) FindByIdNumber(id string) (user *userModel.User, err error) {
	return s.userRepo.FindByIdNumber(id)
}

func (s *UserService) CreateUser(id_number, email, password, nationality string,
	birthSex userModel.BirthSex, genderIdentity userModel.GenderIdentity,
	birthDate datatypes.Date) (user *userModel.User, err error) {
	newUser := userModel.User{
		Password:       password,
		IDNumber:       id_number,
		Email:          email,
		BirthSex:       birthSex,
		GenderIdentity: genderIdentity,
		BirthDate:      birthDate,
		Nationality:    nationality,
	}

	return s.userRepo.CreateUser(newUser)
}
