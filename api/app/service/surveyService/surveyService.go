package surveyService

import (
	"dou-survey/app/model/surveyModel"
	"dou-survey/app/model/voteModel"
	"dou-survey/app/repository/surveyRepository"
)

//SurveyServiceInterface define the survey service interface methods
type SurveyServiceInterface interface {
	Vote(userID, choiceID uint) (vote *voteModel.Vote, err error)
	List(limit, offset uint) (survey []surveyModel.Survey, err error)
	ListWithDetails(limit, offset uint) (survey []surveyModel.Survey, err error)
	FindByID(userId uint) (survey *surveyModel.Survey, err error)
	FindByIDDetailed(userId uint) (survey *surveyModel.Survey, err error)
	Create(create *surveyModel.Survey) (survey *surveyModel.Survey, err error)
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

func (s *SurveyService) Vote(userID, choiceID uint) (vote *voteModel.Vote, err error) {
	return s.surveyRepo.Vote(userID, choiceID)
}

func (s *SurveyService) List(limit, offset uint) (survey []surveyModel.Survey, err error) {
	return s.surveyRepo.List(limit, offset)
}

func (s *SurveyService) ListWithDetails(limit, offset uint) (survey []surveyModel.Survey, err error) {
	return s.surveyRepo.ListWithDetails(limit, offset)
}

func (s *SurveyService) FindByID(userId uint) (survey *surveyModel.Survey, err error) {
	return s.surveyRepo.FindByID(userId)
}

func (s *SurveyService) FindByIDDetailed(userId uint) (survey *surveyModel.Survey, err error) {
	return s.surveyRepo.FindByIDDetailed(userId)
}

func (s *SurveyService) Create(create *surveyModel.Survey) (survey *surveyModel.Survey, err error) {
	return s.surveyRepo.CreateSurvey(create)
}
