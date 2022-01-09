package questionModel

import (
	"dou-survey/app/model/choiceModel"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	SurveyRefer uint
	Choices     []choiceModel.Choice `gorm:"foreignKey:QuestionRefer"`
}
