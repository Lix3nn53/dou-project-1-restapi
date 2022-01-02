package employeeService

import (
	"dou-survey/app/model/employeeModel"
	"dou-survey/app/repository/employeeRepository"
)

//EmployeeServiceInterface define the employee service interface methods
type EmployeeServiceInterface interface {
	FindByUserId(userId uint) (employee *employeeModel.Employee, err error)
}

// billingService handles communication with the employee repository
type EmployeeService struct {
	employeeRepo employeeRepository.EmployeeRepositoryInterface
}

// NewEmployeeService implements the employee service interface.
func NewEmployeeService(employeeRepo employeeRepository.EmployeeRepositoryInterface) EmployeeServiceInterface {
	return &EmployeeService{
		employeeRepo,
	}
}

func (s *EmployeeService) FindByUserId(userId uint) (employee *employeeModel.Employee, err error) {
	return s.employeeRepo.FindByUserId(userId)
}
