package model

import (
	"gorm.io/gorm"
)

type Choice struct {
	gorm.Model    `fake:"skip"`
	QuestionRefer uint   `binding:"required" fake:"skip"`
	Value         string `binding:"required" fake:"{quote}"`
	Votes         []Vote `gorm:"foreignKey:ChoiceRefer" fakesize:"20"`
}
