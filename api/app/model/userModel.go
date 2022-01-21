package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents user resources.
type User struct {
	gorm.Model
	Password       string         `binding:"required"`
	IDNumber       string         `gorm:"unique;not null;column:id_number" binding:"required" valid:"stringlength(11|11)~IDNumber must be 11 digits long"`
	Email          string         `binding:"required" valid:"email"`
	Name           string         `binding:"required"`
	Surname        string         `binding:"required"`
	BirthSex       BirthSex       `binding:"required"`
	GenderIdentity GenderIdentity `binding:"required"`
	BirthDate      time.Time      `binding:"required"`
	IsResident     bool           `binding:"required"`
	Sessions       string
	Employee       Employee `gorm:"foreignKey:UserRefer"`
	Votes          []Vote   `gorm:"foreignKey:UserRefer"`
	CreatedSurveys []Survey `gorm:"foreignKey:UserRefer"`
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

type UserReduced struct {
	IDNumber       string
	Email          string
	Name           string
	Surname        string
	BirthSex       BirthSex
	GenderIdentity GenderIdentity
	BirthDate      time.Time
	IsResident     bool
}
