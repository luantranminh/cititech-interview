package member

import (
	"context"

	"github.com/luantranminh/team-management-app/models"
)

// Repository .
type Repository interface {
	Create(ctx context.Context, name, phone string) (models.Member, error)
	AssignToProject(ctx context.Context, memberID models.UUID, projectID models.UUID) bool
}
