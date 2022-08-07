// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/google/uuid"
	"github.com/thanishsid/goserver/domain"
	"sync"
)

// Ensure, that SessionServiceMock does implement domain.SessionService.
// If this is not the case, regenerate this file with moq.
var _ domain.SessionService = &SessionServiceMock{}

// SessionServiceMock is a mock implementation of domain.SessionService.
//
//	func TestSomethingThatUsesSessionService(t *testing.T) {
//
//		// make and configure a mocked domain.SessionService
//		mockedSessionService := &SessionServiceMock{
//			AddDataFunc: func(ctx context.Context, id domain.SID, key string, value any) error {
//				panic("mock out the AddData method")
//			},
//			ClearDataFunc: func(ctx context.Context, id domain.SID) error {
//				panic("mock out the ClearData method")
//			},
//			CreateFunc: func(ctx context.Context, input domain.CreateSessionInput) (*domain.Session, error) {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(ctx context.Context, id domain.SID) error {
//				panic("mock out the Delete method")
//			},
//			DeleteAllByUserIDFunc: func(ctx context.Context, userID uuid.UUID) error {
//				panic("mock out the DeleteAllByUserID method")
//			},
//			GetFunc: func(ctx context.Context, id domain.SID, userAgent string) (*domain.Session, error) {
//				panic("mock out the Get method")
//			},
//			GetAllByUserIDFunc: func(ctx context.Context, userID uuid.UUID) ([]*domain.Session, error) {
//				panic("mock out the GetAllByUserID method")
//			},
//			RemoveDataFunc: func(ctx context.Context, id domain.SID, key string) error {
//				panic("mock out the RemoveData method")
//			},
//			UpdateRoleByUserIDFunc: func(ctx context.Context, userID uuid.UUID, role domain.Role) error {
//				panic("mock out the UpdateRoleByUserID method")
//			},
//		}
//
//		// use mockedSessionService in code that requires domain.SessionService
//		// and then make assertions.
//
//	}
type SessionServiceMock struct {
	// AddDataFunc mocks the AddData method.
	AddDataFunc func(ctx context.Context, id domain.SID, key string, value any) error

	// ClearDataFunc mocks the ClearData method.
	ClearDataFunc func(ctx context.Context, id domain.SID) error

	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, input domain.CreateSessionInput) (*domain.Session, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, id domain.SID) error

	// DeleteAllByUserIDFunc mocks the DeleteAllByUserID method.
	DeleteAllByUserIDFunc func(ctx context.Context, userID uuid.UUID) error

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, id domain.SID, userAgent string) (*domain.Session, error)

	// GetAllByUserIDFunc mocks the GetAllByUserID method.
	GetAllByUserIDFunc func(ctx context.Context, userID uuid.UUID) ([]*domain.Session, error)

	// RemoveDataFunc mocks the RemoveData method.
	RemoveDataFunc func(ctx context.Context, id domain.SID, key string) error

	// UpdateRoleByUserIDFunc mocks the UpdateRoleByUserID method.
	UpdateRoleByUserIDFunc func(ctx context.Context, userID uuid.UUID, role domain.Role) error

	// calls tracks calls to the methods.
	calls struct {
		// AddData holds details about calls to the AddData method.
		AddData []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID domain.SID
			// Key is the key argument value.
			Key string
			// Value is the value argument value.
			Value any
		}
		// ClearData holds details about calls to the ClearData method.
		ClearData []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID domain.SID
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Input is the input argument value.
			Input domain.CreateSessionInput
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID domain.SID
		}
		// DeleteAllByUserID holds details about calls to the DeleteAllByUserID method.
		DeleteAllByUserID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID domain.SID
			// UserAgent is the userAgent argument value.
			UserAgent string
		}
		// GetAllByUserID holds details about calls to the GetAllByUserID method.
		GetAllByUserID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
		}
		// RemoveData holds details about calls to the RemoveData method.
		RemoveData []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID domain.SID
			// Key is the key argument value.
			Key string
		}
		// UpdateRoleByUserID holds details about calls to the UpdateRoleByUserID method.
		UpdateRoleByUserID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
			// Role is the role argument value.
			Role domain.Role
		}
	}
	lockAddData            sync.RWMutex
	lockClearData          sync.RWMutex
	lockCreate             sync.RWMutex
	lockDelete             sync.RWMutex
	lockDeleteAllByUserID  sync.RWMutex
	lockGet                sync.RWMutex
	lockGetAllByUserID     sync.RWMutex
	lockRemoveData         sync.RWMutex
	lockUpdateRoleByUserID sync.RWMutex
}

// AddData calls AddDataFunc.
func (mock *SessionServiceMock) AddData(ctx context.Context, id domain.SID, key string, value any) error {
	if mock.AddDataFunc == nil {
		panic("SessionServiceMock.AddDataFunc: method is nil but SessionService.AddData was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		ID    domain.SID
		Key   string
		Value any
	}{
		Ctx:   ctx,
		ID:    id,
		Key:   key,
		Value: value,
	}
	mock.lockAddData.Lock()
	mock.calls.AddData = append(mock.calls.AddData, callInfo)
	mock.lockAddData.Unlock()
	return mock.AddDataFunc(ctx, id, key, value)
}

// AddDataCalls gets all the calls that were made to AddData.
// Check the length with:
//
//	len(mockedSessionService.AddDataCalls())
func (mock *SessionServiceMock) AddDataCalls() []struct {
	Ctx   context.Context
	ID    domain.SID
	Key   string
	Value any
} {
	var calls []struct {
		Ctx   context.Context
		ID    domain.SID
		Key   string
		Value any
	}
	mock.lockAddData.RLock()
	calls = mock.calls.AddData
	mock.lockAddData.RUnlock()
	return calls
}

// ClearData calls ClearDataFunc.
func (mock *SessionServiceMock) ClearData(ctx context.Context, id domain.SID) error {
	if mock.ClearDataFunc == nil {
		panic("SessionServiceMock.ClearDataFunc: method is nil but SessionService.ClearData was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  domain.SID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockClearData.Lock()
	mock.calls.ClearData = append(mock.calls.ClearData, callInfo)
	mock.lockClearData.Unlock()
	return mock.ClearDataFunc(ctx, id)
}

// ClearDataCalls gets all the calls that were made to ClearData.
// Check the length with:
//
//	len(mockedSessionService.ClearDataCalls())
func (mock *SessionServiceMock) ClearDataCalls() []struct {
	Ctx context.Context
	ID  domain.SID
} {
	var calls []struct {
		Ctx context.Context
		ID  domain.SID
	}
	mock.lockClearData.RLock()
	calls = mock.calls.ClearData
	mock.lockClearData.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *SessionServiceMock) Create(ctx context.Context, input domain.CreateSessionInput) (*domain.Session, error) {
	if mock.CreateFunc == nil {
		panic("SessionServiceMock.CreateFunc: method is nil but SessionService.Create was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Input domain.CreateSessionInput
	}{
		Ctx:   ctx,
		Input: input,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, input)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedSessionService.CreateCalls())
func (mock *SessionServiceMock) CreateCalls() []struct {
	Ctx   context.Context
	Input domain.CreateSessionInput
} {
	var calls []struct {
		Ctx   context.Context
		Input domain.CreateSessionInput
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *SessionServiceMock) Delete(ctx context.Context, id domain.SID) error {
	if mock.DeleteFunc == nil {
		panic("SessionServiceMock.DeleteFunc: method is nil but SessionService.Delete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  domain.SID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedSessionService.DeleteCalls())
func (mock *SessionServiceMock) DeleteCalls() []struct {
	Ctx context.Context
	ID  domain.SID
} {
	var calls []struct {
		Ctx context.Context
		ID  domain.SID
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// DeleteAllByUserID calls DeleteAllByUserIDFunc.
func (mock *SessionServiceMock) DeleteAllByUserID(ctx context.Context, userID uuid.UUID) error {
	if mock.DeleteAllByUserIDFunc == nil {
		panic("SessionServiceMock.DeleteAllByUserIDFunc: method is nil but SessionService.DeleteAllByUserID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockDeleteAllByUserID.Lock()
	mock.calls.DeleteAllByUserID = append(mock.calls.DeleteAllByUserID, callInfo)
	mock.lockDeleteAllByUserID.Unlock()
	return mock.DeleteAllByUserIDFunc(ctx, userID)
}

// DeleteAllByUserIDCalls gets all the calls that were made to DeleteAllByUserID.
// Check the length with:
//
//	len(mockedSessionService.DeleteAllByUserIDCalls())
func (mock *SessionServiceMock) DeleteAllByUserIDCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
	}
	mock.lockDeleteAllByUserID.RLock()
	calls = mock.calls.DeleteAllByUserID
	mock.lockDeleteAllByUserID.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *SessionServiceMock) Get(ctx context.Context, id domain.SID, userAgent string) (*domain.Session, error) {
	if mock.GetFunc == nil {
		panic("SessionServiceMock.GetFunc: method is nil but SessionService.Get was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		ID        domain.SID
		UserAgent string
	}{
		Ctx:       ctx,
		ID:        id,
		UserAgent: userAgent,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, id, userAgent)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedSessionService.GetCalls())
func (mock *SessionServiceMock) GetCalls() []struct {
	Ctx       context.Context
	ID        domain.SID
	UserAgent string
} {
	var calls []struct {
		Ctx       context.Context
		ID        domain.SID
		UserAgent string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetAllByUserID calls GetAllByUserIDFunc.
func (mock *SessionServiceMock) GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Session, error) {
	if mock.GetAllByUserIDFunc == nil {
		panic("SessionServiceMock.GetAllByUserIDFunc: method is nil but SessionService.GetAllByUserID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockGetAllByUserID.Lock()
	mock.calls.GetAllByUserID = append(mock.calls.GetAllByUserID, callInfo)
	mock.lockGetAllByUserID.Unlock()
	return mock.GetAllByUserIDFunc(ctx, userID)
}

// GetAllByUserIDCalls gets all the calls that were made to GetAllByUserID.
// Check the length with:
//
//	len(mockedSessionService.GetAllByUserIDCalls())
func (mock *SessionServiceMock) GetAllByUserIDCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
	}
	mock.lockGetAllByUserID.RLock()
	calls = mock.calls.GetAllByUserID
	mock.lockGetAllByUserID.RUnlock()
	return calls
}

// RemoveData calls RemoveDataFunc.
func (mock *SessionServiceMock) RemoveData(ctx context.Context, id domain.SID, key string) error {
	if mock.RemoveDataFunc == nil {
		panic("SessionServiceMock.RemoveDataFunc: method is nil but SessionService.RemoveData was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  domain.SID
		Key string
	}{
		Ctx: ctx,
		ID:  id,
		Key: key,
	}
	mock.lockRemoveData.Lock()
	mock.calls.RemoveData = append(mock.calls.RemoveData, callInfo)
	mock.lockRemoveData.Unlock()
	return mock.RemoveDataFunc(ctx, id, key)
}

// RemoveDataCalls gets all the calls that were made to RemoveData.
// Check the length with:
//
//	len(mockedSessionService.RemoveDataCalls())
func (mock *SessionServiceMock) RemoveDataCalls() []struct {
	Ctx context.Context
	ID  domain.SID
	Key string
} {
	var calls []struct {
		Ctx context.Context
		ID  domain.SID
		Key string
	}
	mock.lockRemoveData.RLock()
	calls = mock.calls.RemoveData
	mock.lockRemoveData.RUnlock()
	return calls
}

// UpdateRoleByUserID calls UpdateRoleByUserIDFunc.
func (mock *SessionServiceMock) UpdateRoleByUserID(ctx context.Context, userID uuid.UUID, role domain.Role) error {
	if mock.UpdateRoleByUserIDFunc == nil {
		panic("SessionServiceMock.UpdateRoleByUserIDFunc: method is nil but SessionService.UpdateRoleByUserID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
		Role   domain.Role
	}{
		Ctx:    ctx,
		UserID: userID,
		Role:   role,
	}
	mock.lockUpdateRoleByUserID.Lock()
	mock.calls.UpdateRoleByUserID = append(mock.calls.UpdateRoleByUserID, callInfo)
	mock.lockUpdateRoleByUserID.Unlock()
	return mock.UpdateRoleByUserIDFunc(ctx, userID, role)
}

// UpdateRoleByUserIDCalls gets all the calls that were made to UpdateRoleByUserID.
// Check the length with:
//
//	len(mockedSessionService.UpdateRoleByUserIDCalls())
func (mock *SessionServiceMock) UpdateRoleByUserIDCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
	Role   domain.Role
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
		Role   domain.Role
	}
	mock.lockUpdateRoleByUserID.RLock()
	calls = mock.calls.UpdateRoleByUserID
	mock.lockUpdateRoleByUserID.RUnlock()
	return calls
}