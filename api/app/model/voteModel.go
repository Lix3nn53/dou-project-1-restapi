package model

import "gorm.io/gorm"

type Vote struct {
	gorm.Model  `fake:"skip"`
	UserRefer   uint `gorm:"index:idx_vote,unique" binding:"required" fake:"{number:1,10}"`
	ChoiceRefer uint `gorm:"index:idx_vote,unique" binding:"required" fake:"skip"`
}
