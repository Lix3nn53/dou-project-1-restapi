package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents user resources.
type User struct {
	gorm.Model     `fake:"skip"`
	Password       string         `binding:"required"`
	IDNumber       string         `gorm:"unique;not null;column:id_number" binding:"required" valid:"stringlength(11|11)~IDNumber must be 11 digits long" fake:"{digitn:11}"`
	Email          string         `binding:"required" valid:"email" fake:"{email}"`
	Name           string         `binding:"required" fake:"{firstname}"`
	Surname        string         `binding:"required" fake:"{lastname}"`
	BirthSex       BirthSex       `binding:"required" fake:"{randomstring:[woman,man,norespond]}"`
	GenderIdentity GenderIdentity `binding:"required" fake:"{randomstring:[woman,man,transgender,non,norespond]}"`
	BirthDate      time.Time      `binding:"required" fake:"{daterange:1940-01-01,2020-12-31}" format:"yyyy-MM-dd"`
	IsResident     *bool          `binding:"required"`
	Sessions       string         `fake:"skip"`
	Employee       Employee       `gorm:"foreignKey:UserRefer" fake:"skip"`
	Votes          []Vote         `gorm:"foreignKey:UserRefer" fake:"skip"`
	CreatedSurveys []Survey       `gorm:"foreignKey:UserRefer" fake:"skip"`
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
