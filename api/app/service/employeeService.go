package service

import (
	"dou-survey/app/model"
	"dou-survey/app/repository"
)

//EmployeeServiceInterface define the employee service interface methods
type EmployeeServiceInterface interface {
	FindByUserId(userId uint) (employee *model.Employee, err error)
}

// billingService handles communication with the employee repository
type EmployeeService struct {
	employeeRepo repository.EmployeeRepositoryInterface
}

// NewEmployeeService implements the employee service interface.
func NewEmployeeService(employeeRepo repository.EmployeeRepositoryInterface) EmployeeServiceInterface {
	return &EmployeeService{
		employeeRepo,
	}
}

func (s *EmployeeService) FindByUserId(userId uint) (employee *model.Employee, err error) {
	return s.employeeRepo.FindByUserId(userId)
}
