package repository

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/luantranminh/team-management-app/models"
	"github.com/luantranminh/team-management-app/project"
	"gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// projectRepo .
type projectRepo struct {
	DB *gorm.DB
}

// NewProjectRepository .
func NewProjectRepository(db *gorm.DB) project.Repository {
	return &projectRepo{
		DB: db,
	}
}

// Create create new project with given name
func (m *projectRepo) Create(ctx context.Context, name string) (*models.Project, error) {
	project := models.Project{Name: name}
	return &project, m.DB.Create(&project).Error
}

// GetByID return project name with given id and it's members
func (m *projectRepo) GetByID(ctx context.Context, id models.UUID) (*models.Project, []models.Member, error) {
	project := models.Project{}

	err := validate.Var(id.String(), "uuid4")
	if err != nil {
		return nil, nil, err
	}

	if err := m.DB.Where("id = ?", id).Find(&project).Error; err != nil {
		return nil, nil, err
	}

	members, err := m.getProjectMembers(id)

	return &project, members, err

}

func (m *projectRepo) getProjectMembers(projectID models.UUID) ([]models.Member, error) {
	members := []models.Member{}

	err := m.DB.Table("members").
		Joins("INNER JOIN assignments on assignments.member_id = members.id WHERE assignments.project_id = ?", projectID).
		Scan(&members).Error

	return members, err
}
