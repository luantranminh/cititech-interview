package usecase

import (
	"context"

	"github.com/luantranminh/team-management-app/member"
	"github.com/luantranminh/team-management-app/models"
)

type memberUsecase struct {
	memberRepos member.Repository
}

// NewMemberUsecase .
func NewMemberUsecase(m member.Repository) member.Usecase {
	return &memberUsecase{
		memberRepos: m,
	}
}

// Create function  present for create a user in business layer
func (m memberUsecase) Create(ctx context.Context, name, phone string) (*models.Member, error) {
	return m.memberRepos.Create(ctx, name, phone)
}

// AssignToProject function  present for action 'assign a member to project' in business layer
func (m memberUsecase) AssignToProject(ctx context.Context, memberID models.UUID, projectID models.UUID) error {
	return m.memberRepos.AssignToProject(ctx, memberID, projectID)
}
