package project

import (
	"context"

	"github.com/luantranminh/team-management-app/models"
)

// Usecase .
type Usecase interface {
	Create(ctx context.Context, name string) (*models.Project, error)
	GetByID(ctx context.Context, id models.UUID) (*models.Project, []models.Member, error)
}
