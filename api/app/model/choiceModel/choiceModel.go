package choiceModel

import (
	"gorm.io/gorm"

	"dou-survey/app/model/voteModel"
)

type Choice struct {
	gorm.Model
	QuestionRefer uint
	Value         string
	Votes         []voteModel.Vote `gorm:"foreignKey:ChoiceRefer"`
}
