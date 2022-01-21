package service

import (
	"dou-survey/app/model"
	"dou-survey/app/repository"
)

//SurveyServiceInterface define the survey service interface methods
type SurveyServiceInterface interface {
	ChoiceVotersInfo(choiceID uint) (voters []model.UserReduced, err error)
	Vote(userID, surveyID uint, votes []uint) (created []model.Vote, err error)
	VotedAlready(userID, surveyID uint) (voted bool, err error)
	ListActive(limit, offset uint) (survey []model.Survey, err error)
	ListResults(limit, offset uint) (survey []model.Survey, err error)
	CountActive() (count int, err error)
	CountResults() (count int, err error)
	FindByIDReduced(userId uint) (survey *model.Survey, err error)
	FindByIDWithVotes(userId uint) (survey *model.Survey, err error)
	FindByIDWithoutVotes(userId uint) (survey *model.Survey, err error)
	CountQuestion(id uint) (count int, err error)
	Create(create *model.Survey) (survey *model.Survey, err error)
}

// billingService handles communication with the survey repository
type SurveyService struct {
	surveyRepo repository.SurveyRepositoryInterface
}

// NewSurveyService implements the survey service interface.
func NewSurveyService(surveyRepo repository.SurveyRepositoryInterface) SurveyServiceInterface {
	return &SurveyService{
		surveyRepo,
	}
}

func (s *SurveyService) ChoiceVotersInfo(choiceID uint) (voters []model.UserReduced, err error) {
	return s.surveyRepo.ChoiceVotersInfo(choiceID)
}

func (s *SurveyService) Vote(userID, surveyID uint, votes []uint) (created []model.Vote, err error) {
	return s.surveyRepo.Vote(userID, surveyID, votes)
}

func (s *SurveyService) VotedAlready(userID, surveyID uint) (voted bool, err error) {
	return s.surveyRepo.VotedAlready(userID, surveyID)
}

func (s *SurveyService) CountActive() (count int, err error) {
	return s.surveyRepo.CountActive()
}

func (s *SurveyService) CountResults() (count int, err error) {
	return s.surveyRepo.CountResults()
}

func (s *SurveyService) ListActive(limit, offset uint) (survey []model.Survey, err error) {
	return s.surveyRepo.ListActive(limit, offset)
}

func (s *SurveyService) ListResults(limit, offset uint) (survey []model.Survey, err error) {
	return s.surveyRepo.ListResults(limit, offset)
}

func (s *SurveyService) FindByIDReduced(userId uint) (survey *model.Survey, err error) {
	return s.surveyRepo.FindByIDReduced(userId)
}

func (s *SurveyService) FindByIDWithVotes(userId uint) (survey *model.Survey, err error) {
	return s.surveyRepo.FindByIDWithVotes(userId)
}

func (s *SurveyService) FindByIDWithoutVotes(userId uint) (survey *model.Survey, err error) {
	return s.surveyRepo.FindByIDWithoutVotes(userId)
}

func (s *SurveyService) CountQuestion(id uint) (count int, err error) {
	return s.surveyRepo.CountQuestion(id)
}

func (s *SurveyService) Create(create *model.Survey) (survey *model.Survey, err error) {
	return s.surveyRepo.CreateSurvey(create)
}
