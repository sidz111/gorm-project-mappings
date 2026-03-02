package service

import (
	"context"
	"fmt"

	"github.com/sidz111/employee-mapping-gorm/model"
	"github.com/sidz111/employee-mapping-gorm/repository"
)

type EmployeeService interface {
	Create(ctx context.Context, employee *model.Employee) error
	GetByID(ctx context.Context, id int) (*model.Employee, error)
	GetAll(ctx context.Context) ([]model.Employee, error)
	Update(ctx context.Context, employee *model.Employee) error
	Delete(ctx context.Context, id int) error
}

type employeeService struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{employeeRepo: employeeRepo}
}

func (s *employeeService) Create(ctx context.Context, employee *model.Employee) error {
	if err := validateEmployee(employee); err != nil {
		return err
	}
	return s.employeeRepo.Create(ctx, employee)
}
func (s *employeeService) GetByID(ctx context.Context, id int) (*model.Employee, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid employee ID")
	}
	return s.employeeRepo.GetByID(ctx, id)
}
func (s *employeeService) GetAll(ctx context.Context) ([]model.Employee, error) {
	return s.employeeRepo.GetAll(ctx)
}
func (s *employeeService) Update(ctx context.Context, employee *model.Employee) error {
	if err := validateEmployee(employee); err != nil {
		return err
	}
	return s.employeeRepo.Update(ctx, employee)
}
func (s *employeeService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid employee ID")
	}
	return s.employeeRepo.Delete(ctx, id)
}

func validateEmployee(employee *model.Employee) error {
	if employee.Name == "" {
		return fmt.Errorf("employee name is required")
	}
	if employee.Salary <= 0 {
		return fmt.Errorf("employee salary must be greater than 0")
	}
	if employee.Project.ID == 0 {
		return fmt.Errorf("employee project ID is required")
	}
	return nil
}
