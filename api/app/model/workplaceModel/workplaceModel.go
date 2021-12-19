package workplaceModel

import (
	"goa-golang/app/model/employeeModel"

	"gorm.io/gorm"
)

type Workplace struct {
	gorm.Model
	Name      string
	Type      WorkplaceType
	Employees []employeeModel.Employee `gorm:"foreignKey:WorkplaceRefer"`
}

type WorkplaceType int

const (
	Institution  WorkplaceType = iota // Institution = 0
	Municipality                      // Municipality = 1
	Company                           // Company = 2
	University                        // University = 3
)

func (wp WorkplaceType) String() string {
	return []string{"Institution", "Municipality", "Company", "University"}[wp]
}
