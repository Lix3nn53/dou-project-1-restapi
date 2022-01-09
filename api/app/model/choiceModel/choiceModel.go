package choiceModel

import (
	"gorm.io/gorm"

	"dou-survey/app/model/voteModel"
)

type Choice struct {
	gorm.Model
	QuestionRefer uint             `binding:"required"`
	Value         string           `binding:"required"`
	Votes         []voteModel.Vote `gorm:"foreignKey:ChoiceRefer"`
}
