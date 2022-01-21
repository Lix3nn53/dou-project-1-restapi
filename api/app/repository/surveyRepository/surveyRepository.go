package surveyRepository

import (
	"database/sql"
	"dou-survey/app/model/choiceModel"
	"dou-survey/app/model/questionModel"
	"dou-survey/app/model/surveyModel"
	"dou-survey/app/model/userModel"
	"dou-survey/app/model/voteModel"
	"dou-survey/internal/logger"
	"dou-survey/internal/storage"
	"errors"
)

// billingRepository handles communication with the survey store
type SurveyRepository struct {
	db     *storage.DbStore
	logger logger.Logger
}

//SurveyRepositoryInterface define the survey repository interface methods
type SurveyRepositoryInterface interface {
	ChoiceVotersInfo(choiceID uint) (voters []userModel.UserReduced, err error)
	Vote(userID, surveyID uint, votes []uint) (created []voteModel.Vote, err error)
	VotedAlready(userID, surveyID uint) (voted bool, err error)
	ListActive(limit, offset uint) (surveys []surveyModel.Survey, err error)
	ListResults(limit, offset uint) (surveys []surveyModel.Survey, err error)
	CountActive() (count int, err error)
	CountResults() (count int, err error)
	FindByIDReduced(id uint) (survey *surveyModel.Survey, err error)
	FindByIDWithVotes(id uint) (survey *surveyModel.Survey, err error)
	FindByIDWithoutVotes(id uint) (survey *surveyModel.Survey, err error)
	CountQuestion(id uint) (count int, err error)
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
func (r *SurveyRepository) ChoiceVotersInfo(choiceID uint) (voters []userModel.UserReduced, err error) {
	rows, err := r.db.Raw("SELECT v.id AS vote_id, u.id AS user_id, u.birth_sex, u.gender_identity, u.birth_date, u.nationality FROM (SELECT * FROM `choices` WHERE `choices`.`id` = ?) AS c JOIN votes AS v ON v.choice_refer = c.id JOIN users AS u ON u.id = v.user_refer ORDER BY c.id", choiceID).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// Values to load into
	voters = make([]userModel.UserReduced, 0)

	for rows.Next() {
		var voteID sql.NullInt64
		var userID sql.NullInt64
		user := &userModel.UserReduced{}

		err = rows.Scan(&voteID, &userID, &user.BirthSex, &user.GenderIdentity, &user.BirthDate, &user.Nationality)
		if err != nil {
			return nil, err
		}

		voters = append(voters, *user)
	}

	return voters, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) Vote(userID, surveyID uint, votes []uint) (created []voteModel.Vote, err error) {
	rows, err := r.db.Raw("SELECT q.id as question_id, c.id AS choice_id FROM (SELECT * FROM `surveys` WHERE `surveys`.`id` = ?) AS s JOIN questions AS q ON q.survey_refer = s.id JOIN choices AS c ON c.question_refer = q.id ORDER BY c.id", surveyID).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// Values to load into
	questions := make([]questionModel.Question, 0)

	for rows.Next() {
		var questionID sql.NullInt64
		var choiceID sql.NullInt64

		err = rows.Scan(&questionID, &choiceID)
		if err != nil {
			return nil, err
		}
		if !choiceID.Valid || !questionID.Valid {
			continue
		}

		questionIDInt := uint(questionID.Int64)

		// Check if surveys exists in result array
		var question *questionModel.Question
		isNewQuestion := true
		for i, s := range questions {
			if s.ID == questionIDInt {
				question = &questions[i]
				isNewQuestion = false
			}
		}

		if isNewQuestion {
			question = &questionModel.Question{}
			question.ID = questionIDInt
			question.Choices = make([]choiceModel.Choice, 0)
			questions = append(questions, *question)
			question = &questions[len(questions)-1]
		}

		choiceIDInt := uint(choiceID.Int64)
		choice := choiceModel.Choice{}
		choice.ID = choiceIDInt

		question.Choices = append(question.Choices, choice)
	}

	created = make([]voteModel.Vote, 0)

	// loop for validation
	for i, s := range questions {
		voteIndex := int(votes[i])
		if len(s.Choices) <= voteIndex {
			return nil, errors.New("vote index out of bounds")
		}

		vote := voteModel.Vote{
			UserRefer:   userID,
			ChoiceRefer: s.Choices[voteIndex].ID,
		}

		created = append(created, vote)
	}

	// create
	r.logger.Infof("VOTE")
	r.logger.Infof("%#v", created)
	result := r.db.Create(&created)
	r.logger.Infof("VOTE AFTER CREATE")
	r.logger.Infof("%#v", created)

	if err = result.Error; err != nil {
		return nil, err
	}

	return created, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) VotedAlready(userID, surveyID uint) (voted bool, err error) {
	rows, err := r.db.Raw("SELECT v.id AS vote_id FROM (SELECT * FROM `surveys` WHERE `surveys`.`id` = ?) AS s JOIN questions AS q ON q.survey_refer = s.id JOIN choices AS c ON c.question_refer = q.id LEFT JOIN votes AS v ON v.choice_refer = c.id WHERE v.user_refer = ? ORDER BY c.id", surveyID, userID).Rows()
	if err != nil {
		return true, err
	}

	defer rows.Close()
	voted = false

	for rows.Next() {
		voted = true
		break
	}

	return voted, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) CountActive() (count int, err error) {
	rows, err := r.db.Raw("SELECT count(1) FROM `surveys` AS s WHERE `s`.`deleted_at` IS NULL AND date('now') BETWEEN `s`.`date_start` AND `s`.`date_end`").Rows()
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return -1, err
		}
	}

	return count, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) CountResults() (count int, err error) {
	rows, err := r.db.Raw("SELECT count(1) FROM `surveys` WHERE `surveys`.`deleted_at` IS NULL AND NOT date('now') BETWEEN `surveys`.`date_start` AND `surveys`.`date_end`").Rows()
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return -1, err
		}
	}

	return count, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) ListActive(limit, offset uint) (surveys []surveyModel.Survey, err error) {
	rows, err := r.db.Raw("SELECT s.id, s.user_refer, s.subject, s.description, s.date_start, s.date_end FROM `surveys` AS s WHERE `s`.`deleted_at` IS NULL AND date('now') BETWEEN `s`.`date_start` AND `s`.`date_end` ORDER BY `s`.`id` LIMIT ? OFFSET ?", limit, offset).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// Values to load into
	surveys = make([]surveyModel.Survey, 0)

	for rows.Next() {
		survey := &surveyModel.Survey{}

		err = rows.Scan(&survey.ID, &survey.UserRefer, &survey.Subject, &survey.Description, &survey.DateStart, &survey.DateEnd)
		if err != nil {
			return nil, err
		}

		surveys = append(surveys, *survey)
	}

	return surveys, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) ListResults(limit, offset uint) (surveys []surveyModel.Survey, err error) {
	// Query with joins
	rows, err := r.db.Raw("SELECT s.id, s.user_refer, s.subject, s.description, s.date_start, s.date_end, q.id AS question_id, q.value AS question_value, c.id AS choice_id, c.value AS choice_value, v.id AS vote_id FROM (SELECT * FROM `surveys` WHERE `surveys`.`deleted_at` IS NULL AND NOT date('now') BETWEEN `surveys`.`date_start` AND `surveys`.`date_end` ORDER BY `surveys`.`id` LIMIT ? OFFSET ?) AS s JOIN questions AS q ON q.survey_refer = s.id JOIN choices AS c ON c.question_refer = q.id LEFT JOIN votes AS v ON v.choice_refer = c.id ORDER BY c.id", limit, offset).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// Values to load into
	surveys = make([]surveyModel.Survey, 0)

	for rows.Next() {
		survey := &surveyModel.Survey{}
		survey.Questions = make([]questionModel.Question, 0)

		question := &questionModel.Question{}
		question.Choices = make([]choiceModel.Choice, 0)

		choice := &choiceModel.Choice{}
		choice.Votes = make([]voteModel.Vote, 0)

		var voteID sql.NullInt64

		err = rows.Scan(&survey.ID, &survey.UserRefer, &survey.Subject, &survey.Description, &survey.DateStart, &survey.DateEnd,
			&question.ID, &question.Value, &choice.ID, &choice.Value, &voteID)
		if err != nil {
			return nil, err
		}

		// Check if surveys exists in result array
		isNewSurvey := true
		for i, s := range surveys {
			if s.ID == survey.ID {
				survey = &surveys[i]
				isNewSurvey = false
			}
		}

		// Check if question exists in survey
		isNewQuestion := true
		for i, s := range survey.Questions {
			if s.ID == question.ID {
				question = &survey.Questions[i]
				isNewQuestion = false
			}
		}

		// Check if choice exists in
		isNewChoice := true
		for i, s := range question.Choices {
			if s.ID == choice.ID {
				choice = &question.Choices[i]
				isNewChoice = false
			}
		}

		if !voteID.Valid {
			// vote id is null meaning this choice does not have any votes yet
			// if this choice is added to results already we can ignore this vote
			// but if this choice is not added to results we will add with empty vote array
			if isNewChoice {
				// vote is null but choice is new so lets add this
				question.Choices = append(question.Choices, *choice)
			} // else vote is null and choice is already added so we can ignore this

			if isNewQuestion {
				survey.Questions = append(survey.Questions, *question)
			}

			if isNewSurvey {
				surveys = append(surveys, *survey)
			}
		} else { // vote is not null
			voteID := uint(voteID.Int64)
			vote := voteModel.Vote{}
			vote.ID = voteID
			vote.Model.ID = voteID

			isNewVote := true
			for _, s := range choice.Votes {
				if s.ID == voteID {
					isNewVote = false
				}
			}

			if !isNewVote { // should not be possible, panic to see if it occurs
				r.logger.DPanic("DUBLICATE VOTE")
				continue
			}

			choice.Votes = append(choice.Votes, vote)

			if isNewChoice {
				question.Choices = append(question.Choices, *choice)
			}

			if isNewQuestion {
				survey.Questions = append(survey.Questions, *question)
			}

			if isNewSurvey {
				surveys = append(surveys, *survey)
			}
		}
	}

	return surveys, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) FindByIDReduced(id uint) (survey *surveyModel.Survey, err error) {
	// Query with joins
	rows, err := r.db.Raw("SELECT s.id, s.user_refer, s.subject, s.description, s.date_start, s.date_end FROM `surveys` AS s WHERE `s`.`id` = ?", id).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// Values to load into
	survey = &surveyModel.Survey{}

	for rows.Next() {
		err = rows.Scan(&survey.ID, &survey.UserRefer, &survey.Subject, &survey.Description, &survey.DateStart, &survey.DateEnd)
		if err != nil {
			return nil, err
		}
	}

	return survey, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) FindByIDWithVotes(id uint) (survey *surveyModel.Survey, err error) {
	// Query with joins
	rows, err := r.db.Raw("SELECT s.id, s.user_refer, s.subject, s.description, s.date_start, s.date_end, q.id AS question_id, q.value AS question_value, c.id AS choice_id, c.value AS choice_value, v.id AS vote_id FROM (SELECT * FROM `surveys` WHERE `surveys`.`id` = ?) AS s JOIN questions AS q ON q.survey_refer = s.id JOIN choices AS c ON c.question_refer = q.id LEFT JOIN votes AS v ON v.choice_refer = c.id ORDER BY c.id", id).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// Values to load into
	survey = &surveyModel.Survey{}
	survey.Questions = make([]questionModel.Question, 0)

	for rows.Next() {
		question := &questionModel.Question{}
		question.Choices = make([]choiceModel.Choice, 0)

		choice := &choiceModel.Choice{}
		choice.Votes = make([]voteModel.Vote, 0)

		var voteID sql.NullInt64

		err = rows.Scan(&survey.ID, &survey.UserRefer, &survey.Subject, &survey.Description, &survey.DateStart, &survey.DateEnd,
			&question.ID, &question.Value, &choice.ID, &choice.Value, &voteID)
		if err != nil {
			return nil, err
		}

		// Check if question exists in survey
		isNewQuestion := true
		for i, s := range survey.Questions {
			if s.ID == question.ID {
				question = &survey.Questions[i]
				isNewQuestion = false
			}
		}

		// Check if choice exists in
		isNewChoice := true
		for i, s := range question.Choices {
			if s.ID == choice.ID {
				choice = &question.Choices[i]
				isNewChoice = false
			}
		}

		if !voteID.Valid {
			// vote id is null meaning this choice does not have any votes yet
			// if this choice is added to results already we can ignore this vote
			// but if this choice is not added to results we will add with empty vote array
			if isNewChoice {
				// vote is null but choice is new so lets add this
				question.Choices = append(question.Choices, *choice)
			} // else vote is null and choice is already added so we can ignore this

			if isNewQuestion {
				survey.Questions = append(survey.Questions, *question)
			}
		} else { // vote is not null
			voteID := uint(voteID.Int64)
			vote := voteModel.Vote{}
			vote.ID = voteID
			vote.Model.ID = voteID

			isNewVote := true
			for _, s := range choice.Votes {
				if s.ID == voteID {
					isNewVote = false
				}
			}

			if !isNewVote { // should not be possible, panic to see if it occurs
				r.logger.DPanic("DUBLICATE VOTE")
				continue
			}

			choice.Votes = append(choice.Votes, vote)

			if isNewChoice {
				question.Choices = append(question.Choices, *choice)
			}

			if isNewQuestion {
				survey.Questions = append(survey.Questions, *question)
			}
		}
	}

	return survey, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) FindByIDWithoutVotes(id uint) (survey *surveyModel.Survey, err error) {
	// Query with joins
	rows, err := r.db.Raw("SELECT s.id, s.user_refer, s.subject, s.description, s.date_start, s.date_end, q.id AS question_id, q.value AS question_value, c.id AS choice_id, c.value AS choice_value FROM (SELECT * FROM `surveys` WHERE `surveys`.`id` = ?) AS s JOIN questions AS q ON q.survey_refer = s.id JOIN choices AS c ON c.question_refer = q.id ORDER BY c.id", id).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// Values to load into
	survey = &surveyModel.Survey{}
	survey.Questions = make([]questionModel.Question, 0)

	for rows.Next() {
		question := &questionModel.Question{}
		question.Choices = make([]choiceModel.Choice, 0)

		choice := &choiceModel.Choice{}
		choice.Votes = make([]voteModel.Vote, 0)

		err = rows.Scan(&survey.ID, &survey.UserRefer, &survey.Subject, &survey.Description, &survey.DateStart, &survey.DateEnd,
			&question.ID, &question.Value, &choice.ID, &choice.Value)
		if err != nil {
			return nil, err
		}

		// Check if question exists in survey
		isNewQuestion := true
		for i, s := range survey.Questions {
			if s.ID == question.ID {
				question = &survey.Questions[i]
				isNewQuestion = false
			}
		}

		// Check if choice exists in
		isNewChoice := true
		for i, s := range question.Choices {
			if s.ID == choice.ID {
				choice = &question.Choices[i]
				isNewChoice = false
			}
		}

		// vote id is null meaning this choice does not have any votes yet
		// if this choice is added to results already we can ignore this vote
		// but if this choice is not added to results we will add with empty vote array
		if isNewChoice {
			// vote is null but choice is new so lets add this
			question.Choices = append(question.Choices, *choice)
		} // else vote is null and choice is already added so we can ignore this

		if isNewQuestion {
			survey.Questions = append(survey.Questions, *question)
		}
	}

	return survey, nil
}

// FindByID implements the method to find a survey from the store
func (r *SurveyRepository) CountQuestion(id uint) (count int, err error) {
	// Query with joins
	rows, err := r.db.Raw("SELECT q.id AS question_id FROM (SELECT * FROM `surveys` WHERE `surveys`.`id` = ?) AS s JOIN questions AS q ON q.survey_refer = s.id", id).Rows()
	if err != nil {
		return 0, err
	}

	defer rows.Close()
	// Values to load into
	questions := make([]uint, 0)

	for rows.Next() {
		var questionID uint

		err = rows.Scan(&questionID)
		if err != nil {
			return 0, err
		}

		// Check if question exists in survey
		isNewQuestion := true
		for _, s := range questions {
			if s == questionID {
				isNewQuestion = false
			}
		}

		if isNewQuestion {
			questions = append(questions, questionID)
		}
	}

	return len(questions), nil
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
