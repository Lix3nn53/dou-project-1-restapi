package surveyModel

import (
	"goa-golang/app/model/choiceModel"
	"goa-golang/app/model/voteModel"

	"gorm.io/gorm"
)

type Survey struct {
	gorm.Model
	Choices []choiceModel.Choice `gorm:"foreignKey:SurveyRefer"`
	Votes   []voteModel.Vote     `gorm:"foreignKey:SurveyRefer"`
}
