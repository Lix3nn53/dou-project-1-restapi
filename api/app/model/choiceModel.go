package model

import (
	"gorm.io/gorm"
)

type Choice struct {
	gorm.Model
	QuestionRefer uint   `binding:"required"`
	Value         string `binding:"required"`
	Votes         []Vote `gorm:"foreignKey:ChoiceRefer"`
}
