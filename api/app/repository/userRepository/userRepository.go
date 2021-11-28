package userRepository

import (
	"goa-golang/app/model/userModel"
	"goa-golang/internal/storage"
	"strings"
)

// billingRepository handles communication with the user store
type UserRepository struct {
	db *storage.DbStore
}

//UserRepositoryInterface define the user repository interface methods
type UserRepositoryInterface interface {
	FindByID(id string) (user *userModel.User, err error)
	FindByIdNumber(id string) (user *userModel.User, err error)
	RemoveByID(uuid string) error
	UpdateByID(id string, user userModel.User) error
	CreateUser(create userModel.User) (user *userModel.User, err error)
	GetSessions(id string) (sessions string, err error)
	AddSession(id string, refreshToken string) error
	RemoveSession(id string, refreshToken string) error
}

// NewUserRepository implements the user repository interface.
func NewUserRepository(db *storage.DbStore) UserRepositoryInterface {
	return &UserRepository{
		db,
	}
}

// FindByID implements the method to find a user from the store
func (r *UserRepository) FindByID(id string) (user *userModel.User, err error) {
	result := r.db.First(&user, id)

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
func (r *UserRepository) RemoveByID(id string) error {
	result := r.db.Delete(&userModel.User{}, id)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

// UpdateByID implements the method to update a user into the store
func (r *UserRepository) UpdateByID(id string, userUpdate userModel.User) error {
	var user userModel.User

	result := r.db.First(&user, id).Updates(userUpdate)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

// Create implements the method to persist a new user
func (r *UserRepository) CreateUser(userCreate userModel.User) (user *userModel.User, err error) {
	result := r.db.Create(&user) // pass pointer of data to Create

	if err = result.Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindByID implements the method to find a user from the store
func (r *UserRepository) GetSessions(id string) (sessions string, err error) {
	var user userModel.User

	result := r.db.Select("sessions").First(&user, id)

	if err := result.Error; err != nil {
		return "", err
	}

	return user.Sessions, nil
}

func (r *UserRepository) AddSession(id string, refreshToken string) error {
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

func (r *UserRepository) RemoveSession(id string, refreshToken string) error {
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
