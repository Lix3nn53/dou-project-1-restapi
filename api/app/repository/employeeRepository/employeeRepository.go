package employeeRepository

import (
	"goa-golang/app/model/employeeModel"
	"goa-golang/internal/storage"
)

// billingRepository handles communication with the employee store
type EmployeeRepository struct {
	db *storage.DbStore
}

//EmployeeRepositoryInterface define the employee repository interface methods
type EmployeeRepositoryInterface interface {
	FindByID(id string) (employee *employeeModel.Employee, err error)
	FindByUserId(userId string) (employee *employeeModel.Employee, err error)
	RemoveByID(uuid string) error
	UpdateByID(id string, employee employeeModel.Employee) error
	CreateEmployee(create employeeModel.Employee) (employee *employeeModel.Employee, err error)
}

// NewEmployeeRepository implements the employee repository interface.
func NewEmployeeRepository(db *storage.DbStore) EmployeeRepositoryInterface {
	return &EmployeeRepository{
		db,
	}
}

// FindByID implements the method to find a employee from the store
func (r *EmployeeRepository) FindByID(id string) (employee *employeeModel.Employee, err error) {
	result := r.db.First(&employee, id)

	if err := result.Error; err != nil {
		return nil, err
	}

	return employee, nil
}

// FindByUserId implements the method to find a employee from the store
func (r *EmployeeRepository) FindByUserId(userId string) (employee *employeeModel.Employee, err error) {
	result := r.db.First(employee, "UserRefer = ?", userId)

	if err := result.Error; err != nil {
		return nil, err
	}

	return employee, nil
}

// RemoveByID implements the method to remove a employee from the store
func (r *EmployeeRepository) RemoveByID(id string) error {
	result := r.db.Delete(&employeeModel.Employee{}, id)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

// UpdateByID implements the method to update a employee into the store
func (r *EmployeeRepository) UpdateByID(id string, employeeUpdate employeeModel.Employee) error {
	var employee employeeModel.Employee

	result := r.db.First(&employee, id).Updates(employeeUpdate)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}

// Create implements the method to persist a new employee
func (r *EmployeeRepository) CreateEmployee(employeeCreate employeeModel.Employee) (_ *employeeModel.Employee, err error) {
	result := r.db.Create(&employeeCreate) // pass pointer of data to Create

	if err = result.Error; err != nil {
		return nil, err
	}

	return &employeeCreate, nil
}
