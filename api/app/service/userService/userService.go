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
	CreateUser(*CreateUserDTO) (user *userModel.User, err error)
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

type CreateUserDTO struct {
	IDNumber       string                   `json:"IDNumber" binding:"required" valid:"stringlength(11|11)"`
	Email          string                   `json:"email" binding:"required" valid:"email"`
	Password       string                   `json:"password" binding:"required"`
	Name           string                   `json:"name" binding:"required"`
	Surname        string                   `json:"surname" binding:"required"`
	Nationality    string                   `json:"nationality" binding:"required"`
	BirthSex       userModel.BirthSex       `json:"birthSex" binding:"required"`
	GenderIdentity userModel.GenderIdentity `json:"genderIdentity" binding:"required"`
	BirthDate      datatypes.Date           `json:"birthDate" binding:"required"`
}

func (s *UserService) CreateUser(dto *CreateUserDTO) (user *userModel.User, err error) {
	newUser := userModel.User{
		Password:       dto.Password,
		IDNumber:       dto.IDNumber,
		Email:          dto.Email,
		Name:           dto.Name,
		Surname:        dto.Surname,
		BirthSex:       dto.BirthSex,
		GenderIdentity: dto.GenderIdentity,
		BirthDate:      dto.BirthDate,
		Nationality:    dto.Nationality,
	}

	return s.userRepo.CreateUser(newUser)
}
