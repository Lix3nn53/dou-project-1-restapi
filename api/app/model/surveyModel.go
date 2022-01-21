package model

import (
	"time"

	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model
	UserRefer   uint
	Questions   []Question `gorm:"foreignKey:SurveyRefer" binding:"required" valid:"required~You must add at least one question"`
	Subject     string     `binding:"required"`
	Description string     `binding:"required"`
	DateStart   time.Time  `binding:"required"`
	DateEnd     time.Time  `binding:"required"`
}
