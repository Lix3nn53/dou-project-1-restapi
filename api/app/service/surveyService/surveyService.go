package surveyService

import (
	"dou-survey/app/model/surveyModel"
	"dou-survey/app/repository/surveyRepository"
)

//SurveyServiceInterface define the survey service interface methods
type SurveyServiceInterface interface {
	Vote(userID, surveyID uint, votes []uint) (err error)
	VotedAlready(userID, surveyID uint) (voted bool, err error)
	List(limit, offset uint) (survey []surveyModel.Survey, err error)
	ListWithDetails(limit, offset uint) (survey []surveyModel.Survey, err error)
	FindByIDReduced(userId uint) (survey *surveyModel.Survey, err error)
	FindByIDWithVotes(userId uint) (survey *surveyModel.Survey, err error)
	FindByIDWithoutVotes(userId uint) (survey *surveyModel.Survey, err error)
	CountChoice(id uint) (count uint, err error)
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

func (s *SurveyService) Vote(userID, surveyID uint, votes []uint) (err error) {
	return s.surveyRepo.Vote(userID, surveyID, votes)
}

func (s *SurveyService) VotedAlready(userID, surveyID uint) (voted bool, err error) {
	return s.surveyRepo.VotedAlready(userID, surveyID)
}

func (s *SurveyService) List(limit, offset uint) (survey []surveyModel.Survey, err error) {
	return s.surveyRepo.List(limit, offset)
}

func (s *SurveyService) ListWithDetails(limit, offset uint) (survey []surveyModel.Survey, err error) {
	return s.surveyRepo.ListWithDetails(limit, offset)
}

func (s *SurveyService) FindByIDReduced(userId uint) (survey *surveyModel.Survey, err error) {
	return s.surveyRepo.FindByIDReduced(userId)
}

func (s *SurveyService) FindByIDWithVotes(userId uint) (survey *surveyModel.Survey, err error) {
	return s.surveyRepo.FindByIDWithVotes(userId)
}

func (s *SurveyService) FindByIDWithoutVotes(userId uint) (survey *surveyModel.Survey, err error) {
	return s.surveyRepo.FindByIDWithoutVotes(userId)
}

func (s *SurveyService) CountChoice(id uint) (count uint, err error) {
	return s.surveyRepo.CountChoice(id)
}

func (s *SurveyService) Create(create *surveyModel.Survey) (survey *surveyModel.Survey, err error) {
	return s.surveyRepo.CreateSurvey(create)
}
