package userModel

import "gorm.io/gorm"

// User represents user resources.
type User struct {
	gorm.Model
	IdNumber string `json:"id_number" db:"id_number"`
	Email    string `json:"email" db:"email"`
	Sessions string `json:"sessions" db:"sessions"`
}
