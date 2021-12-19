package userModel

import (
	"goa-golang/app/model/documentModel"
	"goa-golang/app/model/employeeModel"

	"gorm.io/gorm"
)

// User represents user resources.
type User struct {
	gorm.Model
	Password  string                   `json:"password" db:"password"`
	TCKN      string                   `json:"id_number" db:"id_number"  gorm:"unique;not null"`
	Email     string                   `json:"email" db:"email"`
	Sessions  string                   `json:"sessions" db:"sessions"`
	Documents []documentModel.Document `gorm:"foreignKey:UserRefer"`
	Employee  employeeModel.Employee   `gorm:"foreignKey:UserRefer"`
}
