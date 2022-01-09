package userModel

import (
	"dou-survey/app/model/employeeModel"
	"dou-survey/app/model/voteModel"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// User represents user resources.
type User struct {
	gorm.Model
	Password       string                 `json:"password" binding:"required"`
	IDNumber       string                 `json:"IDNumber" gorm:"unique;not null;column:id_number" binding:"required" valid:"stringlength(11|11)"`
	Email          string                 `json:"email" binding:"required" valid:"email"`
	Name           string                 `json:"name" binding:"required"`
	Surname        string                 `json:"surname" binding:"required"`
	BirthSex       BirthSex               `json:"birthSex" binding:"required"`
	GenderIdentity GenderIdentity         `json:"genderIdentity" binding:"required"`
	BirthDate      datatypes.Date         `json:"birthDate" binding:"required"`
	Nationality    string                 `json:"nationality" binding:"required"`
	Sessions       string                 `json:"sessions"`
	Employee       employeeModel.Employee `gorm:"foreignKey:UserRefer"`
	Votes          []voteModel.Vote       `gorm:"foreignKey:UserRefer"`
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
