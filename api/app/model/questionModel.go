package model

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	SurveyRefer uint     `binding:"required"`
	Value       string   `binding:"required"`
	Choices     []Choice `gorm:"foreignKey:QuestionRefer" binding:"required"`
}
