// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"sync"

	"github.com/luantranminh/team-management-app/models"
)

var (
	lockUsecaseMockCreate  sync.RWMutex
	lockUsecaseMockGetByID sync.RWMutex
)

// UsecaseMock is a mock implementation of Usecase.
//
//     func TestSomethingThatUsesUsecase(t *testing.T) {
//
//         // make and configure a mocked Usecase
//         mockedUsecase := &UsecaseMock{
//             CreateFunc: func(ctx context.Context, name string) (*models.Project, error) {
// 	               panic("TODO: mock out the Create method")
//             },
//             GetByIDFunc: func(ctx context.Context, id models.UUID) (*models.Project, []models.Member, error) {
// 	               panic("TODO: mock out the GetByID method")
//             },
//         }
//
//         // TODO: use mockedUsecase in code that requires Usecase
//         //       and then make assertions.
//
//     }
type UsecaseMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, name string) (*models.Project, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, id models.UUID) (*models.Project, []models.Member, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID models.UUID
		}
	}
}

// Create calls CreateFunc.
func (mock *UsecaseMock) Create(ctx context.Context, name string) (*models.Project, error) {
	if mock.CreateFunc == nil {
		panic("UsecaseMock.CreateFunc: method is nil but Usecase.Create was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	lockUsecaseMockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	lockUsecaseMockCreate.Unlock()
	return mock.CreateFunc(ctx, name)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedUsecase.CreateCalls())
func (mock *UsecaseMock) CreateCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	lockUsecaseMockCreate.RLock()
	calls = mock.calls.Create
	lockUsecaseMockCreate.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *UsecaseMock) GetByID(ctx context.Context, id models.UUID) (*models.Project, []models.Member, error) {
	if mock.GetByIDFunc == nil {
		panic("UsecaseMock.GetByIDFunc: method is nil but Usecase.GetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  models.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockUsecaseMockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	lockUsecaseMockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, id)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//     len(mockedUsecase.GetByIDCalls())
func (mock *UsecaseMock) GetByIDCalls() []struct {
	Ctx context.Context
	ID  models.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  models.UUID
	}
	lockUsecaseMockGetByID.RLock()
	calls = mock.calls.GetByID
	lockUsecaseMockGetByID.RUnlock()
	return calls
}