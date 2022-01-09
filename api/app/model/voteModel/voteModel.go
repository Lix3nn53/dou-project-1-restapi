package voteModel

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	UserRefer   uint `gorm:"primaryKey"`
	ChoiceRefer uint `gorm:"primaryKey"`
}
