// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/google/uuid"
	"github.com/thanishsid/goserver/domain"
	"sync"
)

// Ensure, that UserServiceMock does implement domain.UserService.
// If this is not the case, regenerate this file with moq.
var _ domain.UserService = &UserServiceMock{}

// UserServiceMock is a mock implementation of domain.UserService.
//
//	func TestSomethingThatUsesUserService(t *testing.T) {
//
//		// make and configure a mocked domain.UserService
//		mockedUserService := &UserServiceMock{
//			AllByIDSFunc: func(ctx context.Context, ids ...uuid.UUID) ([]*domain.User, error) {
//				panic("mock out the AllByIDS method")
//			},
//			ChangeRoleFunc: func(ctx context.Context, input domain.RoleChangeInput) error {
//				panic("mock out the ChangeRole method")
//			},
//			CompleteRegistrationFunc: func(ctx context.Context, input domain.CompleteRegistrationInput) (*domain.User, error) {
//				panic("mock out the CompleteRegistration method")
//			},
//			CreateFunc: func(ctx context.Context, input domain.CreateUserInput) (*domain.User, error) {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(ctx context.Context, id uuid.UUID) error {
//				panic("mock out the Delete method")
//			},
//			InitRegistrationFunc: func(ctx context.Context, input domain.InitRegistrationInput) error {
//				panic("mock out the InitRegistration method")
//			},
//			ManyFunc: func(ctx context.Context, filter domain.UserFilterInput) (*domain.ListData[domain.User], error) {
//				panic("mock out the Many method")
//			},
//			OneFunc: func(ctx context.Context, id uuid.UUID) (*domain.User, error) {
//				panic("mock out the One method")
//			},
//			OneByEmailFunc: func(ctx context.Context, email string) (*domain.User, error) {
//				panic("mock out the OneByEmail method")
//			},
//			UpdateFunc: func(ctx context.Context, userID uuid.UUID, input domain.UserUpdateInput) error {
//				panic("mock out the Update method")
//			},
//		}
//
//		// use mockedUserService in code that requires domain.UserService
//		// and then make assertions.
//
//	}
type UserServiceMock struct {
	// AllByIDSFunc mocks the AllByIDS method.
	AllByIDSFunc func(ctx context.Context, ids ...uuid.UUID) ([]*domain.User, error)

	// ChangeRoleFunc mocks the ChangeRole method.
	ChangeRoleFunc func(ctx context.Context, input domain.RoleChangeInput) error

	// CompleteRegistrationFunc mocks the CompleteRegistration method.
	CompleteRegistrationFunc func(ctx context.Context, input domain.CompleteRegistrationInput) (*domain.User, error)

	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, input domain.CreateUserInput) (*domain.User, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, id uuid.UUID) error

	// InitRegistrationFunc mocks the InitRegistration method.
	InitRegistrationFunc func(ctx context.Context, input domain.InitRegistrationInput) error

	// ManyFunc mocks the Many method.
	ManyFunc func(ctx context.Context, filter domain.UserFilterInput) (*domain.ListData[domain.User], error)

	// OneFunc mocks the One method.
	OneFunc func(ctx context.Context, id uuid.UUID) (*domain.User, error)

	// OneByEmailFunc mocks the OneByEmail method.
	OneByEmailFunc func(ctx context.Context, email string) (*domain.User, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, userID uuid.UUID, input domain.UserUpdateInput) error

	// calls tracks calls to the methods.
	calls struct {
		// AllByIDS holds details about calls to the AllByIDS method.
		AllByIDS []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Ids is the ids argument value.
			Ids []uuid.UUID
		}
		// ChangeRole holds details about calls to the ChangeRole method.
		ChangeRole []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Input is the input argument value.
			Input domain.RoleChangeInput
		}
		// CompleteRegistration holds details about calls to the CompleteRegistration method.
		CompleteRegistration []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Input is the input argument value.
			Input domain.CompleteRegistrationInput
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Input is the input argument value.
			Input domain.CreateUserInput
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// InitRegistration holds details about calls to the InitRegistration method.
		InitRegistration []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Input is the input argument value.
			Input domain.InitRegistrationInput
		}
		// Many holds details about calls to the Many method.
		Many []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter domain.UserFilterInput
		}
		// One holds details about calls to the One method.
		One []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// OneByEmail holds details about calls to the OneByEmail method.
		OneByEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Email is the email argument value.
			Email string
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
			// Input is the input argument value.
			Input domain.UserUpdateInput
		}
	}
	lockAllByIDS             sync.RWMutex
	lockChangeRole           sync.RWMutex
	lockCompleteRegistration sync.RWMutex
	lockCreate               sync.RWMutex
	lockDelete               sync.RWMutex
	lockInitRegistration     sync.RWMutex
	lockMany                 sync.RWMutex
	lockOne                  sync.RWMutex
	lockOneByEmail           sync.RWMutex
	lockUpdate               sync.RWMutex
}

// AllByIDS calls AllByIDSFunc.
func (mock *UserServiceMock) AllByIDS(ctx context.Context, ids ...uuid.UUID) ([]*domain.User, error) {
	if mock.AllByIDSFunc == nil {
		panic("UserServiceMock.AllByIDSFunc: method is nil but UserService.AllByIDS was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Ids []uuid.UUID
	}{
		Ctx: ctx,
		Ids: ids,
	}
	mock.lockAllByIDS.Lock()
	mock.calls.AllByIDS = append(mock.calls.AllByIDS, callInfo)
	mock.lockAllByIDS.Unlock()
	return mock.AllByIDSFunc(ctx, ids...)
}

// AllByIDSCalls gets all the calls that were made to AllByIDS.
// Check the length with:
//
//	len(mockedUserService.AllByIDSCalls())
func (mock *UserServiceMock) AllByIDSCalls() []struct {
	Ctx context.Context
	Ids []uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		Ids []uuid.UUID
	}
	mock.lockAllByIDS.RLock()
	calls = mock.calls.AllByIDS
	mock.lockAllByIDS.RUnlock()
	return calls
}

// ChangeRole calls ChangeRoleFunc.
func (mock *UserServiceMock) ChangeRole(ctx context.Context, input domain.RoleChangeInput) error {
	if mock.ChangeRoleFunc == nil {
		panic("UserServiceMock.ChangeRoleFunc: method is nil but UserService.ChangeRole was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Input domain.RoleChangeInput
	}{
		Ctx:   ctx,
		Input: input,
	}
	mock.lockChangeRole.Lock()
	mock.calls.ChangeRole = append(mock.calls.ChangeRole, callInfo)
	mock.lockChangeRole.Unlock()
	return mock.ChangeRoleFunc(ctx, input)
}

// ChangeRoleCalls gets all the calls that were made to ChangeRole.
// Check the length with:
//
//	len(mockedUserService.ChangeRoleCalls())
func (mock *UserServiceMock) ChangeRoleCalls() []struct {
	Ctx   context.Context
	Input domain.RoleChangeInput
} {
	var calls []struct {
		Ctx   context.Context
		Input domain.RoleChangeInput
	}
	mock.lockChangeRole.RLock()
	calls = mock.calls.ChangeRole
	mock.lockChangeRole.RUnlock()
	return calls
}

// CompleteRegistration calls CompleteRegistrationFunc.
func (mock *UserServiceMock) CompleteRegistration(ctx context.Context, input domain.CompleteRegistrationInput) (*domain.User, error) {
	if mock.CompleteRegistrationFunc == nil {
		panic("UserServiceMock.CompleteRegistrationFunc: method is nil but UserService.CompleteRegistration was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Input domain.CompleteRegistrationInput
	}{
		Ctx:   ctx,
		Input: input,
	}
	mock.lockCompleteRegistration.Lock()
	mock.calls.CompleteRegistration = append(mock.calls.CompleteRegistration, callInfo)
	mock.lockCompleteRegistration.Unlock()
	return mock.CompleteRegistrationFunc(ctx, input)
}

// CompleteRegistrationCalls gets all the calls that were made to CompleteRegistration.
// Check the length with:
//
//	len(mockedUserService.CompleteRegistrationCalls())
func (mock *UserServiceMock) CompleteRegistrationCalls() []struct {
	Ctx   context.Context
	Input domain.CompleteRegistrationInput
} {
	var calls []struct {
		Ctx   context.Context
		Input domain.CompleteRegistrationInput
	}
	mock.lockCompleteRegistration.RLock()
	calls = mock.calls.CompleteRegistration
	mock.lockCompleteRegistration.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *UserServiceMock) Create(ctx context.Context, input domain.CreateUserInput) (*domain.User, error) {
	if mock.CreateFunc == nil {
		panic("UserServiceMock.CreateFunc: method is nil but UserService.Create was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Input domain.CreateUserInput
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
//	len(mockedUserService.CreateCalls())
func (mock *UserServiceMock) CreateCalls() []struct {
	Ctx   context.Context
	Input domain.CreateUserInput
} {
	var calls []struct {
		Ctx   context.Context
		Input domain.CreateUserInput
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *UserServiceMock) Delete(ctx context.Context, id uuid.UUID) error {
	if mock.DeleteFunc == nil {
		panic("UserServiceMock.DeleteFunc: method is nil but UserService.Delete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
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
//	len(mockedUserService.DeleteCalls())
func (mock *UserServiceMock) DeleteCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// InitRegistration calls InitRegistrationFunc.
func (mock *UserServiceMock) InitRegistration(ctx context.Context, input domain.InitRegistrationInput) error {
	if mock.InitRegistrationFunc == nil {
		panic("UserServiceMock.InitRegistrationFunc: method is nil but UserService.InitRegistration was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Input domain.InitRegistrationInput
	}{
		Ctx:   ctx,
		Input: input,
	}
	mock.lockInitRegistration.Lock()
	mock.calls.InitRegistration = append(mock.calls.InitRegistration, callInfo)
	mock.lockInitRegistration.Unlock()
	return mock.InitRegistrationFunc(ctx, input)
}

// InitRegistrationCalls gets all the calls that were made to InitRegistration.
// Check the length with:
//
//	len(mockedUserService.InitRegistrationCalls())
func (mock *UserServiceMock) InitRegistrationCalls() []struct {
	Ctx   context.Context
	Input domain.InitRegistrationInput
} {
	var calls []struct {
		Ctx   context.Context
		Input domain.InitRegistrationInput
	}
	mock.lockInitRegistration.RLock()
	calls = mock.calls.InitRegistration
	mock.lockInitRegistration.RUnlock()
	return calls
}

// Many calls ManyFunc.
func (mock *UserServiceMock) Many(ctx context.Context, filter domain.UserFilterInput) (*domain.ListData[domain.User], error) {
	if mock.ManyFunc == nil {
		panic("UserServiceMock.ManyFunc: method is nil but UserService.Many was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Filter domain.UserFilterInput
	}{
		Ctx:    ctx,
		Filter: filter,
	}
	mock.lockMany.Lock()
	mock.calls.Many = append(mock.calls.Many, callInfo)
	mock.lockMany.Unlock()
	return mock.ManyFunc(ctx, filter)
}

// ManyCalls gets all the calls that were made to Many.
// Check the length with:
//
//	len(mockedUserService.ManyCalls())
func (mock *UserServiceMock) ManyCalls() []struct {
	Ctx    context.Context
	Filter domain.UserFilterInput
} {
	var calls []struct {
		Ctx    context.Context
		Filter domain.UserFilterInput
	}
	mock.lockMany.RLock()
	calls = mock.calls.Many
	mock.lockMany.RUnlock()
	return calls
}

// One calls OneFunc.
func (mock *UserServiceMock) One(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	if mock.OneFunc == nil {
		panic("UserServiceMock.OneFunc: method is nil but UserService.One was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockOne.Lock()
	mock.calls.One = append(mock.calls.One, callInfo)
	mock.lockOne.Unlock()
	return mock.OneFunc(ctx, id)
}

// OneCalls gets all the calls that were made to One.
// Check the length with:
//
//	len(mockedUserService.OneCalls())
func (mock *UserServiceMock) OneCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockOne.RLock()
	calls = mock.calls.One
	mock.lockOne.RUnlock()
	return calls
}

// OneByEmail calls OneByEmailFunc.
func (mock *UserServiceMock) OneByEmail(ctx context.Context, email string) (*domain.User, error) {
	if mock.OneByEmailFunc == nil {
		panic("UserServiceMock.OneByEmailFunc: method is nil but UserService.OneByEmail was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Email string
	}{
		Ctx:   ctx,
		Email: email,
	}
	mock.lockOneByEmail.Lock()
	mock.calls.OneByEmail = append(mock.calls.OneByEmail, callInfo)
	mock.lockOneByEmail.Unlock()
	return mock.OneByEmailFunc(ctx, email)
}

// OneByEmailCalls gets all the calls that were made to OneByEmail.
// Check the length with:
//
//	len(mockedUserService.OneByEmailCalls())
func (mock *UserServiceMock) OneByEmailCalls() []struct {
	Ctx   context.Context
	Email string
} {
	var calls []struct {
		Ctx   context.Context
		Email string
	}
	mock.lockOneByEmail.RLock()
	calls = mock.calls.OneByEmail
	mock.lockOneByEmail.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *UserServiceMock) Update(ctx context.Context, userID uuid.UUID, input domain.UserUpdateInput) error {
	if mock.UpdateFunc == nil {
		panic("UserServiceMock.UpdateFunc: method is nil but UserService.Update was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
		Input  domain.UserUpdateInput
	}{
		Ctx:    ctx,
		UserID: userID,
		Input:  input,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, userID, input)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedUserService.UpdateCalls())
func (mock *UserServiceMock) UpdateCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
	Input  domain.UserUpdateInput
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
		Input  domain.UserUpdateInput
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}