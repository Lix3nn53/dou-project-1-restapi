package employeeModel

import (
	"gorm.io/gorm"
)

// Employee represents employee resources.
type Employee struct {
	gorm.Model
	UserRefer uint `binding:"required"`
}
