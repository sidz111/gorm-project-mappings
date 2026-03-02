package service

import (
	"context"
	"fmt"

	"github.com/sidz111/employee-mapping-gorm/model"
	"github.com/sidz111/employee-mapping-gorm/repository"
)

type ProjectService interface {
	Create(ctx context.Context, project *model.Project) error
	GetByID(ctx context.Context, id int) (*model.Project, error)
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	Update(ctx context.Context, project *model.Project) error
	Delete(ctx context.Context, id int) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &projectService{projectRepo: projectRepo}
}

func (s *projectService) Create(ctx context.Context, project *model.Project) error {
	if err := validateProject(project); err != nil {
		return err
	}
	return s.projectRepo.Create(ctx, project)
}
func (s *projectService) GetByID(ctx context.Context, id int) (*model.Project, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid project ID")
	}
	return s.projectRepo.GetByID(ctx, id)
}
func (s *projectService) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	return s.projectRepo.GetAllProjects(ctx)
}
func (s *projectService) Update(ctx context.Context, project *model.Project) error {
	if err := validateProject(project); err != nil {
		return err
	}
	return s.projectRepo.Update(ctx, project)
}
func (s *projectService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid project ID")
	}
	return s.projectRepo.Delete(ctx, id)
}

func validateProject(project *model.Project) error {
	if project.Title == "" {
		return fmt.Errorf("project title is required")
	}
	return nil
}
