package voteModel

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	UserRefer   uint `gorm:"primaryKey" binding:"required"`
	ChoiceRefer uint `gorm:"primaryKey" binding:"required"`
}
