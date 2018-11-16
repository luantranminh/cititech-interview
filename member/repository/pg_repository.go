package repository

import (
	"context"
	"regexp"

	"github.com/jinzhu/gorm"
	"github.com/luantranminh/team-management-app/member"
	"github.com/luantranminh/team-management-app/models"
	validator "gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var (
	validate       = validator.New()
	vietnameseName = `^[a-zA-Zaáàảãạâấầẩẫậăắằẳẵặeéèẻẽẹêếềểễệiíìỉĩịoóòỏõọôốồổỗộơớờởỡợuúùủũụưứừửữựyýỳỷỹỵđ ]+$`
	nameRegex      = regexp.MustCompile(vietnameseName)
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
func (m *memberRepo) Create(ctx context.Context, name, phone string) (*models.Member, error) {

	if !nameRegex.MatchString(name) {
		return nil, models.ErrorInvalidName
	}

	if err := validate.Var(phone, "numeric,min=9,max=14"); err != nil {
		return nil, models.ErrorInvalidPhone
	}

	member := models.Member{Name: name, Phone: phone}
	return &member, m.DB.Create(&member).Error
}

// AssignToProject .
func (m *memberRepo) AssignToProject(ctx context.Context, memberID models.UUID, projectID models.UUID) error {

	if err := validate.Var(memberID.String(), "uuid4"); err != nil {
		return models.ErrorInvalidID
	}

	if err := validate.Var(projectID.String(), "uuid4"); err != nil {
		return models.ErrorInvalidID
	}

	assignment := models.Assignment{
		MemberID:  memberID,
		ProjectID: projectID,
	}

	return m.DB.Create(&assignment).Error
}
