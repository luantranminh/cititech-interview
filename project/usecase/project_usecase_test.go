package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/luantranminh/team-management-app/project/mocks"

	"github.com/luantranminh/team-management-app/models"
)

func Test_projectUsecase_Create(t *testing.T) {

	mock := &mocks.RepositoryMock{
		CreateFunc: func(ctx context.Context, name string) (*models.Project, error) {
			return nil, nil
		},
		GetByIDFunc: func(ctx context.Context, id models.UUID) (*models.Project, []models.Member, error) {
			return nil, nil, nil
		},
	}

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Project
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := projectUsecase{
				projectRepos: mock,
			}
			got, err := m.Create(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("projectUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectUsecase_GetByID(t *testing.T) {
	type fields struct {
		projectRepos *mocks.RepositoryMock
	}
	type args struct {
		ctx context.Context
		id  models.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Project
		want1   []models.Member
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := projectUsecase{
				projectRepos: tt.fields.projectRepos,
			}
			got, got1, err := m.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectUsecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("projectUsecase.GetByID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("projectUsecase.GetByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
