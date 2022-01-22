package faker

import (
	"dou-survey/app/model"
	"dou-survey/internal/logger"
	"dou-survey/internal/storage"

	"github.com/brianvoe/gofakeit/v6"
)

// billingRepository handles communication with the survey store
type SurveyFaker struct {
	db     *storage.DbStore
	logger logger.Logger
}

//UserRepositoryInterface define the user repository interface methods
type SurveyFakerInterface interface {
	Generate(userID uint) (survey *model.Survey, err error)
}

// NewSurveyRepository implements the survey repository interface.
func NewSurveyFaker(db *storage.DbStore, logger logger.Logger) SurveyFakerInterface {
	return &SurveyFaker{
		db,
		logger,
	}
}

// FindByID implements the method to find a user from the store
func (r *SurveyFaker) Generate(userID uint) (survey *model.Survey, err error) {
	// TODO implement

	var f model.Survey
	err = gofakeit.Struct(&f)
	if err != nil {
		r.logger.Error(err)
		return
	}

	f.UserRefer = userID

	for i, question := range f.Questions {
		for y, choice := range question.Choices {
			votes := removeDuplicateVotes(choice.Votes)
			f.Questions[i].Choices[y].Votes = votes
		}
	}

	// r.logger.Infof("%+v", f)
	r.logger.JSON(f)

	return &f, nil
}

func removeDuplicateVotes(votes []model.Vote) []model.Vote {
	allUsers := make(map[int]bool)
	list := make([]model.Vote, 0)
	for _, item := range votes {
		if _, value := allUsers[int(item.UserRefer)]; !value {
			allUsers[int(item.UserRefer)] = true
			list = append(list, item)
		}
	}
	return list
}
