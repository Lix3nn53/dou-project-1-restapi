package model

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model  `fake:"skip"`
	SurveyRefer uint     `binding:"required" fake:"skip"`
	Value       string   `binding:"required" fake:"{question}"`
	Choices     []Choice `gorm:"foreignKey:QuestionRefer" binding:"required" fakesize:"3"`
}
