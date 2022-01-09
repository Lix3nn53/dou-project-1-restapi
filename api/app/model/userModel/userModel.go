package userModel

import (
	"dou-survey/app/model/employeeModel"
	"dou-survey/app/model/surveyModel"
	"dou-survey/app/model/voteModel"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// User represents user resources.
type User struct {
	gorm.Model
	Password       string         `binding:"required"`
	IDNumber       string         `gorm:"unique;not null;column:id_number" binding:"required" valid:"stringlength(11|11)"`
	Email          string         `binding:"required" valid:"email"`
	Name           string         `binding:"required"`
	Surname        string         `binding:"required"`
	BirthSex       BirthSex       `binding:"required"`
	GenderIdentity GenderIdentity `binding:"required"`
	BirthDate      datatypes.Date `binding:"required"`
	Nationality    string         `binding:"required"`
	Sessions       string
	Employee       employeeModel.Employee `gorm:"foreignKey:UserRefer"`
	Votes          []voteModel.Vote       `gorm:"foreignKey:UserRefer"`
	CreatedSurveys []surveyModel.Survey   `gorm:"foreignKey:UserRefer"`
}

type BirthSex string

const (
	BirthSexWoman     BirthSex = "woman"
	BirthSexMan       BirthSex = "man"
	BirthSexNoRespond BirthSex = "norespond" // prefer not to respond
)

type GenderIdentity string

const (
	GenderIdentityWoman       GenderIdentity = "woman"
	GenderIdentityMan         GenderIdentity = "man"
	GenderIdentityTransgender GenderIdentity = "transgender"
	GenderIdentityNonBin      GenderIdentity = "non" // non-binary/non-conforming
	GenderIdentityNoRespond   GenderIdentity = "norespond"
)
