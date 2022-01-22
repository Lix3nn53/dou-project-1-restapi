package faker

import (
	"dou-survey/app/model"
	"dou-survey/internal/logger"
	"dou-survey/internal/storage"

	"github.com/brianvoe/gofakeit/v6"
)

// billingRepository handles communication with the user store
type UserFaker struct {
	db     *storage.DbStore
	logger logger.Logger
}

//UserRepositoryInterface define the user repository interface methods
type UserFakerInterface interface {
	Generate(amount int) (err error)
}

// NewUserRepository implements the user repository interface.
func NewUserFaker(db *storage.DbStore, logger logger.Logger) UserFakerInterface {
	return &UserFaker{
		db,
		logger,
	}
}

// FindByID implements the method to find a user from the store
func (r *UserFaker) Generate(amount int) (err error) {
	toCreate := make([]model.User, 0)

	for index := 0; index < amount; index++ {
		var f model.User
		err = gofakeit.Struct(&f)
		if err != nil {
			r.logger.Error(err)
			return
		}

		toCreate = append(toCreate, f)

		// r.logger.Infof("%+v", f)
		r.logger.JSON(f)
	}

	result := r.db.CreateInBatches(toCreate, amount)
	if err = result.Error; err != nil {
		r.logger.Error(err)
		return err
	}

	return nil
}
