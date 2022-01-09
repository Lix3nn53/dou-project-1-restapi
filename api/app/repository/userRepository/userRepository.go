package userRepository

import (
	passwordHelper "dou-survey/app/helpers/password"
	"dou-survey/app/model/userModel"
	"dou-survey/internal/storage"
	"strings"
)

// billingRepository handles communication with the user store
type UserRepository struct {
	db *storage.DbStore
}

//UserRepositoryInterface define the user repository interface methods
type UserRepositoryInterface interface {
	FindByID(id uint) (user *userModel.User, err error)
	FindByIDReduced(id uint) (user *userModel.User, err error)
	FindByIdNumber(id string) (user *userModel.User, err error)
	RemoveByID(id uint) error
	UpdateByID(id uint, user userModel.User) error
	CreateUser(create userModel.User) (user *userModel.User, err error)
	GetSessions(id uint) (sessions string, err error)
	AddSession(id uint, refreshToken string) error
	RemoveSession(id uint, refreshToken string) error
}

// NewUserRepository implements the user repository interface.
func NewUserRepository(db *storage.DbStore) UserRepositoryInterface {
	return &UserRepository{
		db,
	}
}

// FindByID implements the method to find a user from the store
func (r *UserRepository) FindByID(id uint) (user *userModel.User, err error) {
	result := r.db.First(&user, id)

	if err := result.Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindByID implements the method to find a user from the store
func (r *UserRepository) FindByIDReduced(id uint) (user *userModel.User, err error) {
	result := r.db.Select("id_number", "email", "name", "surname", "birth_sex", "gender_identity", "birth_date", "nationality").First(&user, id)

	if err := result.Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindByID implements the method to find a user from the store
func (r *UserRepository) FindByIdNumber(id string) (user *userModel.User, err error) {
	result := r.db.First(&user, "id_number = ?", id)

	if err := result.Error; err != nil {
		return nil, err
	}

	return user, nil
}

// RemoveByID implements the method to remove a user from the store
func (r *UserRepository) RemoveByID(id uint) error {
	result := r.db.Delete(&userModel.User{}, id)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

// UpdateByID implements the method to update a user into the store
func (r *UserRepository) UpdateByID(id uint, userUpdate userModel.User) error {
	var user userModel.User

	result := r.db.First(&user, id).Updates(userUpdate)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

// Create implements the method to persist a new user
func (r *UserRepository) CreateUser(userCreate userModel.User) (_ *userModel.User, err error) {
	hashed, err := passwordHelper.HashPassword(userCreate.Password)

	if err != nil {
		return nil, err
	}

	userCreate.Password = hashed

	result := r.db.Create(&userCreate) // pass pointer of data to Create

	if err = result.Error; err != nil {
		return nil, err
	}

	return &userCreate, nil
}

// FindByID implements the method to find a user from the store
func (r *UserRepository) GetSessions(id uint) (sessions string, err error) {
	var user userModel.User

	result := r.db.Select("sessions").First(&user, id)

	if err := result.Error; err != nil {
		return "", err
	}

	return user.Sessions, nil
}

func (r *UserRepository) AddSession(id uint, refreshToken string) error {
	sessionsStr := refreshToken

	sessions, err := r.GetSessions(id)
	if err != nil {
		return err
	}
	sessionsStr = sessions + "/" + refreshToken

	err = r.UpdateByID(id, userModel.User{Sessions: sessionsStr})
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) RemoveSession(id uint, refreshToken string) error {
	sessions, err := r.GetSessions(id)
	if err != nil {
		return err
	}

	sessionsStr := strings.Replace(sessions, "/"+refreshToken, "", -1)

	err = r.UpdateByID(id, userModel.User{Sessions: sessionsStr})
	if err != nil {
		return err
	}

	return nil
}
