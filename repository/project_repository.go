package repository

import (
	"context"
	"fmt"

	"github.com/sidz111/employee-mapping-gorm/model"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *model.Project) error
	GetByID(ctx context.Context, id int) (*model.Project, error)
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	Update(ctx context.Context, project *model.Project) error
	Delete(ctx context.Context, id int) error
}

type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepo{db: db}
}

func (r *projectRepo) Create(ctx context.Context, project *model.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

func (r *projectRepo) GetByID(ctx context.Context, id int) (*model.Project, error) {
	var project model.Project
	err := r.db.WithContext(ctx).Preload("Employee").First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepo) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	var projects []model.Project
	err := r.db.WithContext(ctx).Preload("Employee").Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}
func (r *projectRepo) Update(ctx context.Context, project *model.Project) error {
	result := r.db.WithContext(ctx).Save(project)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("project not found")
	}
	return nil
}
func (r *projectRepo) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&model.Project{}, id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("project not found")
	}
	return nil
}
