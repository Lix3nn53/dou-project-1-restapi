package surveyModel

import (
	"dou-survey/app/model/questionModel"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model
	Questions   []questionModel.Question `gorm:"foreignKey:SurveyRefer"`
	Subject     string
	Description string
	DateStart   datatypes.Date
	DateEnd     datatypes.Date
}
