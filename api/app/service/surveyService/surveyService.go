package surveyService

import (
	"dou-survey/app/model/surveyModel"
	"dou-survey/app/repository/surveyRepository"
)

//SurveyServiceInterface define the survey service interface methods
type SurveyServiceInterface interface {
	FindByID(userId uint) (survey *surveyModel.Survey, err error)
	Create(create surveyModel.Survey) (survey *surveyModel.Survey, err error)
}

// billingService handles communication with the survey repository
type SurveyService struct {
	surveyRepo surveyRepository.SurveyRepositoryInterface
}

// NewSurveyService implements the survey service interface.
func NewSurveyService(surveyRepo surveyRepository.SurveyRepositoryInterface) SurveyServiceInterface {
	return &SurveyService{
		surveyRepo,
	}
}

func (s *SurveyService) FindByID(userId uint) (survey *surveyModel.Survey, err error) {
	return s.surveyRepo.FindByID(userId)
}

func (s *SurveyService) Create(create surveyModel.Survey) (survey *surveyModel.Survey, err error) {
	return s.surveyRepo.CreateSurvey(create)
}
