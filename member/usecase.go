package member

import (
	"context"

	"github.com/luantranminh/team-management-app/models"
)

// Usecase .
type Usecase interface {
	Create(ctx context.Context, name, phone string) (models.Member, error)
	AssignToProject(ctx context.Context, memberID models.UUID, projectID models.UUID) bool
}
