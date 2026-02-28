package repository

import (
	"context"

	"github.com/sidz111/employee-mapping-gorm/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee *model.Employee) error
	GetByID(ctx context.Context, id int) (*model.Employee, error)
	GetAll(ctx context.Context) ([]model.Employee, error)
	Update(ctx context.Context, employee *model.Employee) error
	Delete(ctx context.Context, id int) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(ctx context.Context, employee *model.Employee) error {
	return r.db.WithContext(ctx).Create(employee).Error
}
func (r *employeeRepository) GetByID(ctx context.Context, id int) (*model.Employee, error) {
	var emp model.Employee
	if err := r.db.WithContext(ctx).Preload("Project").First(&emp, id).Error; err != nil {
		return nil, err
	}
	return &emp, nil
}
func (r *employeeRepository) GetAll(ctx context.Context) ([]model.Employee, error) {
	var employees []model.Employee
	if err := r.db.WithContext(ctx).Preload("Project").Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}
func (r *employeeRepository) Update(ctx context.Context, employee *model.Employee) error {
	if err := r.db.WithContext(ctx).Save(employee).Error; err != nil {
		return err
	}
	return nil
}
func (r *employeeRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&model.Employee{}, id).Error; err != nil {
		return err
	}
	return nil
}
