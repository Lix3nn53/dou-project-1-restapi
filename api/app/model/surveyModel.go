package model

import (
	"time"

	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model  `fake:"skip"`
	UserRefer   uint       `fake:"skip"`
	Questions   []Question `gorm:"foreignKey:SurveyRefer" binding:"required" valid:"required~You must add at least one question" fakesize:"2"`
	Subject     string     `binding:"required" fake:"{sentence:3}"`
	Description string     `binding:"required" fake:"{paragraph:2,3,4,aa}"`
	DateStart   time.Time  `binding:"required"`
	DateEnd     time.Time  `binding:"required"`
}
