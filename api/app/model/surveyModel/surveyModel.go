package surveyModel

import (
	"dou-survey/app/model/choiceModel"
	"dou-survey/app/model/voteModel"

	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model
	Choices []choiceModel.Choice `gorm:"foreignKey:SurveyRefer"`
	Votes   []voteModel.Vote     `gorm:"foreignKey:SurveyRefer"`
}
