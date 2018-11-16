package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/luantranminh/team-management-app/member"
	"github.com/luantranminh/team-management-app/models"
)

// MemberRepo .
type memberRepo struct {
	DB *gorm.DB
}

// NewMemberRepository .
func NewMemberRepository(db *gorm.DB) member.Repository {
	return &memberRepo{
		DB: db,
	}
}

// Create create new member with given name and phone
func (m *memberRepo) Create(ctx context.Context, name, phone string) (models.Member, error) {
	member := models.Member{Name: name, Phone: phone}
	return member, m.DB.Create(&member).Error
}

// AssignToProject .
func (m *memberRepo) AssignToProject(ctx context.Context, memberID models.UUID, projectID models.UUID) bool {
	assignment := models.Assignment{
		MemberID:  memberID,
		ProjectID: memberID,
	}

	if err := m.DB.Create(&assignment).Error; err != nil {
		return false
	}

	return true
}
