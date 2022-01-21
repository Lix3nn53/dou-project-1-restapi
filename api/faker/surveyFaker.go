package faker

import (
	"dou-survey/app/model"
	"dou-survey/internal/logger"
	"dou-survey/internal/storage"
)

// billingRepository handles communication with the survey store
type SurveyFaker struct {
	db     *storage.DbStore
	logger logger.Logger
}

//UserRepositoryInterface define the user repository interface methods
type SurveyFakerInterface interface {
	Generate() (survey *model.Survey, err error)
}

// NewSurveyRepository implements the survey repository interface.
func NewSurveyFaker(db *storage.DbStore, logger logger.Logger) SurveyFakerInterface {
	return &SurveyFaker{
		db,
		logger,
	}
}

// FindByID implements the method to find a user from the store
func (r *SurveyFaker) Generate() (survey *model.Survey, err error) {
	// TODO implement

	return survey, nil
}
