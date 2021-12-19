package employeeModel

import (
	"goa-golang/app/model/documentModel"

	"gorm.io/gorm"
)

// Employee represents employee resources.
type Employee struct {
	gorm.Model
	UserRefer      uint
	Documents      []documentModel.Document `gorm:"foreignKey:EmployeeRefer"`
	WorkplaceRefer uint
}
