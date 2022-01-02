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
	Password       string                 `json:"password" db:"password"`
	IDNumber       string                 `json:"id_number" db:"id_number"  gorm:"unique;not null;column:id_number"`
	Email          string                 `json:"email" db:"email"`
	Name           string                 `json:"name" db:"name"`
	Surname        string                 `json:"surname" db:"surname"`
	Age            uint                   `json:"age" db:"age"`
	BirthSex       BirthSex               `json:"birth_sex" db:"birth_sex"`
	GenderIdentity GenderIdentity         `json:"gender" db:"gender"`
	BirthDate      datatypes.Date         `json:"birth_date" db:"birth_date"`
	Nationality    string                 `json:"nationality" db:"nationality"`
	Sessions       string                 `json:"sessions" db:"sessions"`
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
	GenderIdentityNonBinay    GenderIdentity = "non" // non-binary/non-conforming
	GenderIdentityNoRespond   GenderIdentity = "norespond"
)
