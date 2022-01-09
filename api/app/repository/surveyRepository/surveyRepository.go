package surveyRepository

import (
	"dou-survey/app/model/surveyModel"
	"dou-survey/internal/storage"
)

// billingRepository handles communication with the survey store
type SurveyRepository struct {
	db *storage.DbStore
}

//SurveyRepositoryInterface define the survey repository interface methods
type SurveyRepositoryInterface interface {
	List(limit, offset int) (surveys *[]surveyModel.Survey, err error)
	FindByID(id uint) (survey *surveyModel.Survey, err error)
	RemoveByID(id uint) error
	UpdateByID(id uint, survey surveyModel.Survey) error
	CreateSurvey(create *surveyModel.Survey) (survey *surveyModel.Survey, err error)
}

// NewSurveyRepository implements the survey repository interface.
func NewSurveyRepository(db *storage.DbStore) SurveyRepositoryInterface {
	return &SurveyRepository{
		db,
	}
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) List(limit, offset int) (surveys *[]surveyModel.Survey, err error) {
	result := r.db.Limit(limit).Offset(offset).Find(&surveys)

	if err := result.Error; err != nil {
		return nil, err
	}

	return surveys, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) FindByID(id uint) (survey *surveyModel.Survey, err error) {
	result := r.db.First(&survey, id)

	if err := result.Error; err != nil {
		return nil, err
	}

	return survey, nil
}

// RemoveByID implements the method to remove a survey from the store
func (r *SurveyRepository) RemoveByID(id uint) error {
	result := r.db.Delete(&surveyModel.Survey{}, id)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

// UpdateByID implements the method to update a survey into the store
func (r *SurveyRepository) UpdateByID(id uint, surveyUpdate surveyModel.Survey) error {
	var survey surveyModel.Survey

	result := r.db.First(&survey, id).Updates(surveyUpdate)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

// Create implements the method to persist a new survey
func (r *SurveyRepository) CreateSurvey(surveyCreate *surveyModel.Survey) (_ *surveyModel.Survey, err error) {
	result := r.db.Create(&surveyCreate) // pass pointer of data to Create

	if err = result.Error; err != nil {
		return nil, err
	}

	return surveyCreate, nil
}
