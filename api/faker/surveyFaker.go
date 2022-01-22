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
	Generate(amount int) (err error)
}

// NewSurveyRepository implements the survey repository interface.
func NewSurveyFaker(db *storage.DbStore, logger logger.Logger) SurveyFakerInterface {
	return &SurveyFaker{
		db,
		logger,
	}
}

// FindByID implements the method to find a user from the store
func (r *SurveyFaker) Generate(amount int) (err error) {
	toCreate := make([]model.Survey, 0)

	for index := 0; index < amount; index++ {
		var f model.Survey
		err = gofakeit.Struct(&f)
		if err != nil {
			r.logger.Error(err)
			return
		}

		f.UserRefer = 1

		for i, question := range f.Questions {
			for y, choice := range question.Choices {
				votes := removeDuplicateVotes(choice.Votes)
				f.Questions[i].Choices[y].Votes = votes
			}
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
