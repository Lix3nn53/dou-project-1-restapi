package choiceModel

import "gorm.io/gorm"

type Choice struct {
	gorm.Model
	SurveyRefer uint
	Value       string
}
