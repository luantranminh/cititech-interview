package repository

import (
	"context"
	"reflect"
	"testing"

	dbtest "github.com/luantranminh/team-management-app/config/database/pg/util"
	"github.com/luantranminh/team-management-app/models"
)

func Test_projectRepo_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := dbtest.CreateTestDatabase(t)
	defer cleanup()
	err := dbtest.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Project
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name: "project 1",
			},
			wantErr: false,
			want: &models.Project{
				Name: "project 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &projectRepo{
				DB: testDB,
			}
			got, err := m.Create(context.Background(), tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Because id was generated
			tt.want.ID = got.ID
			tt.want.CreatedAt = got.CreatedAt
			tt.want.UpdatedAt = got.UpdatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("projectRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectRepo_GetByID(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := dbtest.CreateTestDatabase(t)
	defer cleanup()
	err := dbtest.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}
	type args struct {
		id models.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Project
		want1   []models.Member
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &projectRepo{
				DB: testDB,
			}
			got, got1, err := m.GetByID(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectRepo.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("projectRepo.GetByID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("projectRepo.GetByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_projectRepo_getProjectMembers(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := dbtest.CreateTestDatabase(t)
	defer cleanup()
	err := dbtest.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}
	type args struct {
		projectID models.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Member
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &projectRepo{
				DB: testDB,
			}
			got, err := m.getProjectMembers(tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectRepo.getProjectMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("projectRepo.getProjectMembers() = %v, want %v", got, tt.want)
			}
		})
	}
}
