package userModel

import (
	"goa-golang/app/model/employeeModel"
	"goa-golang/app/model/voteModel"

	"gorm.io/gorm"
)

// User represents user resources.
type User struct {
	gorm.Model
	Password string                 `json:"password" db:"password"`
	TCKN     string                 `json:"id_number" db:"id_number"  gorm:"unique;not null;column:id_number"`
	Email    string                 `json:"email" db:"email"`
	Sessions string                 `json:"sessions" db:"sessions"`
	Employee employeeModel.Employee `gorm:"foreignKey:UserRefer"`
	Votes    []voteModel.Vote       `gorm:"foreignKey:UserRefer"`
}
