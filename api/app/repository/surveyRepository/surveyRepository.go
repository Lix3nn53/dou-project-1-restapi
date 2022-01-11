package surveyRepository

import (
	"database/sql"
	"dou-survey/app/model/choiceModel"
	"dou-survey/app/model/questionModel"
	"dou-survey/app/model/surveyModel"
	"dou-survey/app/model/voteModel"
	"dou-survey/internal/logger"
	"dou-survey/internal/storage"
)

// billingRepository handles communication with the survey store
type SurveyRepository struct {
	db     *storage.DbStore
	logger logger.Logger
}

//SurveyRepositoryInterface define the survey repository interface methods
type SurveyRepositoryInterface interface {
	List(limit, offset int) (surveys []surveyModel.Survey, err error)
	FindByID(id uint) (survey *surveyModel.Survey, err error)
	RemoveByID(id uint) error
	UpdateByID(id uint, survey surveyModel.Survey) error
	CreateSurvey(create *surveyModel.Survey) (survey *surveyModel.Survey, err error)
}

// NewSurveyRepository implements the survey repository interface.
func NewSurveyRepository(db *storage.DbStore, logger logger.Logger) SurveyRepositoryInterface {
	return &SurveyRepository{
		db,
		logger,
	}
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) List(limit, offset int) (surveys []surveyModel.Survey, err error) {
	// Query with joins
	rows, err := r.db.Raw("SELECT s.id, s.user_refer, s.subject, s.description, s.date_start, s.date_end, q.id AS question_id, q.value AS question_value, c.id AS choice_id, c.value AS choice_value, v.id AS vote_id FROM (SELECT * FROM `surveys` ORDER BY `surveys`.`id` LIMIT ? OFFSET ?) AS s JOIN questions AS q ON q.survey_refer = s.id JOIN choices AS c ON c.question_refer = q.id LEFT JOIN votes AS v ON v.choice_refer = c.id WHERE `s`.`deleted_at` IS NULL", limit, offset).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// Values to load into
	surveys = make([]surveyModel.Survey, 0)

	for rows.Next() {
		survey := surveyModel.Survey{}
		survey.Questions = make([]questionModel.Question, 0)

		question := questionModel.Question{}
		question.Choices = make([]choiceModel.Choice, 0)

		choice := choiceModel.Choice{}
		choice.Votes = make([]voteModel.Vote, 0)

		var voteID sql.NullInt64

		err = rows.Scan(&survey.ID, &survey.UserRefer, &survey.Subject, &survey.Description, &survey.DateStart, &survey.DateEnd,
			&question.ID, &question.Value, &choice.ID, &choice.Value, &voteID)
		if err != nil {
			return nil, err
		}

		// Check if surveys exists in result array
		for _, s := range surveys {
			if s.ID == survey.ID {
				survey = s
			}
		}

		// Check if question exists in survey
		for _, s := range survey.Questions {
			if s.ID == question.ID {
				question = s
			}
		}

		// Check if choice exists in
		isNewChoice := true
		for _, s := range question.Choices {
			if s.ID == choice.ID {
				choice = s
				isNewChoice = false
			}
		}

		if !voteID.Valid {
			// vote id is null meaning this choice does not have any votes yet
			// if this choice is added to results already we can ignore this vote
			// but if this choice is not added to results we will add with empty vote array

			if isNewChoice {
				// vote is null but choice is new so lets add this
				question.Choices = append(question.Choices, choice)
				survey.Questions = append(survey.Questions, question)
				surveys = append(surveys, survey)
			} // else vote is null and choice is already added so we can ignore this
		} else { // vote is not null
			voteValue := uint(voteID.Int64)
			vote := voteModel.Vote{}
			vote.ID = voteValue
			vote.Model.ID = voteValue

			isNewVote := true
			for _, s := range choice.Votes {
				if s.ID == voteValue {
					isNewVote = false
				}
			}

			if !isNewVote { // should not be possible, panic to see if it occurs
				r.logger.DPanic("DUBLICATE VOTE")
				continue
			}

			choice.Votes = append(choice.Votes, vote)
			question.Choices = append(question.Choices, choice)
			survey.Questions = append(survey.Questions, question)
			surveys = append(surveys, survey)
		}
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
