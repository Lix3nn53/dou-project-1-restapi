package model

import (
	"time"

	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model    `fake:"skip"`
	UserRefer     uint          `fake:"skip"`
	Questions     []Question    `gorm:"foreignKey:SurveyRefer" binding:"required" valid:"required~You must add at least one question" fakesize:"5"`
	Subject       string        `binding:"required" fake:"{sentence:3}"`
	Description   string        `binding:"required" fake:"{paragraph:2,3,4,aa}"`
	DateStart     time.Time     `binding:"required" fake:"{daterange:2022-01-01,2022-03-25}" format:"yyyy-MM-dd"`
	DateEnd       time.Time     `binding:"required" fake:"{daterange:2022-02-25,2022-04-25}" format:"yyyy-MM-dd"`
	ConfirmStatus ConfirmStatus `gorm:"default:waiting" fake:"{randomstring:[waiting,declined,confirmed]}"`
}

type ConfirmStatus string

const (
	ConfirmStatusWaiting  ConfirmStatus = "waiting"
	ConfirmStatusDeclined ConfirmStatus = "declined"
	ConfirmStatusAccepted ConfirmStatus = "confirmed"
)
