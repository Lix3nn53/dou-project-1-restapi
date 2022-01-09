package surveyModel

import (
	"dou-survey/app/model/questionModel"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model
	UserRefer   uint                     `binding:"required"`
	Questions   []questionModel.Question `gorm:"foreignKey:SurveyRefer" binding:"required"`
	Subject     string                   `binding:"required"`
	Description string                   `binding:"required"`
	DateStart   datatypes.Date           `binding:"required"`
	DateEnd     datatypes.Date           `binding:"required"`
}
