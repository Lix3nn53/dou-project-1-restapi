package voteModel

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	UserRefer   uint `gorm:"index:idx_vote,unique" binding:"required"`
	ChoiceRefer uint `gorm:"index:idx_vote,unique" binding:"required"`
}
