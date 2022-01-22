package model

import (
	"gorm.io/gorm"
)

type Choice struct {
	gorm.Model    `fake:"skip"`
	QuestionRefer uint   `binding:"required" fake:"skip"`
	Value         string `binding:"required" fake:"{randomstring:[Extremely well,Very well,Somewhat well,Not so well]}"`
	Votes         []Vote `gorm:"foreignKey:ChoiceRefer" fakesize:"100"`
}
