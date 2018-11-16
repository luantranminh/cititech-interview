package usecase

import (
	"context"

	"github.com/luantranminh/team-management-app/models"
	"github.com/luantranminh/team-management-app/project"
)

type projectUsecase struct {
	projectRepos project.Repository
}

// NewProjectUsecase .
func NewProjectUsecase(m project.Repository) project.Usecase {
	return &projectUsecase{
		projectRepos: m,
	}
}

// Create function  present for create a project in business layer
func (m projectUsecase) Create(ctx context.Context, name string) (*models.Project, error) {
	return m.projectRepos.Create(ctx, name)
}

// GetByID return a project
func (m projectUsecase) GetByID(ctx context.Context, id models.UUID) (*models.Project, []models.Member, error) {
	return m.projectRepos.GetByID(ctx, id)
}
