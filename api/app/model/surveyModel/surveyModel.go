package surveyModel

import (
	"dou-survey/app/model/questionModel"
	"time"

	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model
	UserRefer   uint
	Questions   []questionModel.Question `gorm:"foreignKey:SurveyRefer" binding:"required"`
	Subject     string                   `binding:"required"`
	Description string                   `binding:"required"`
	DateStart   time.Time                `binding:"required"`
	DateEnd     time.Time                `binding:"required"`
}
