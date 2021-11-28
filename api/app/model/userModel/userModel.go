package userModel

import (
	"goa-golang/app/model/documentModel"

	"gorm.io/gorm"
)

// User represents user resources.
type User struct {
	gorm.Model
	TCKN      string                   `json:"id_number" db:"id_number"`
	Email     string                   `json:"email" db:"email"`
	Sessions  string                   `json:"sessions" db:"sessions"`
	Documents []documentModel.Document `gorm:"foreignKey:UserRefer"`
}
