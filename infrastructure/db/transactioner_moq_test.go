// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package db

import (
	"context"
	"github.com/google/uuid"
	"sync"
)

// Ensure, that TransactionerMock does implement Transactioner.
// If this is not the case, regenerate this file with moq.
var _ Transactioner = &TransactionerMock{}

// TransactionerMock is a mock implementation of Transactioner.
//
//	func TestSomethingThatUsesTransactioner(t *testing.T) {
//
//		// make and configure a mocked Transactioner
//		mockedTransactioner := &TransactionerMock{
//			CheckImageExistsFunc: func(ctx context.Context, id uuid.UUID) (bool, error) {
//				panic("mock out the CheckImageExists method")
//			},
//			CheckImageHashExistsFunc: func(ctx context.Context, fileHash []byte) (bool, error) {
//				panic("mock out the CheckImageHashExists method")
//			},
//			CommitFunc: func(ctx context.Context) error {
//				panic("mock out the Commit method")
//			},
//			DeleteImageFunc: func(ctx context.Context, id uuid.UUID) error {
//				panic("mock out the DeleteImage method")
//			},
//			DeleteRoleFunc: func(ctx context.Context, id string) error {
//				panic("mock out the DeleteRole method")
//			},
//			GetAllImagesInIDSFunc: func(ctx context.Context, imageIds []uuid.UUID) ([]GetAllImagesInIDSRow, error) {
//				panic("mock out the GetAllImagesInIDS method")
//			},
//			GetAllRolesFunc: func(ctx context.Context) ([]Role, error) {
//				panic("mock out the GetAllRoles method")
//			},
//			GetAllUsersInIDSFunc: func(ctx context.Context, userIds []uuid.UUID) ([]GetAllUsersInIDSRow, error) {
//				panic("mock out the GetAllUsersInIDS method")
//			},
//			GetImageByIdFunc: func(ctx context.Context, id uuid.UUID) (Image, error) {
//				panic("mock out the GetImageById method")
//			},
//			GetManyImagesFunc: func(ctx context.Context, arg GetManyImagesParams) ([]GetManyImagesRow, error) {
//				panic("mock out the GetManyImages method")
//			},
//			GetManyUsersFunc: func(ctx context.Context, arg GetManyUsersParams) ([]GetManyUsersRow, error) {
//				panic("mock out the GetManyUsers method")
//			},
//			GetUserByEmailFunc: func(ctx context.Context, email string) (GetUserByEmailRow, error) {
//				panic("mock out the GetUserByEmail method")
//			},
//			GetUserByIdFunc: func(ctx context.Context, userID uuid.UUID) (GetUserByIdRow, error) {
//				panic("mock out the GetUserById method")
//			},
//			GetUsersCountFunc: func(ctx context.Context) (int64, error) {
//				panic("mock out the GetUsersCount method")
//			},
//			GetUsersCountByRoleFunc: func(ctx context.Context, userRole string) (int64, error) {
//				panic("mock out the GetUsersCountByRole method")
//			},
//			HardDeleteUserFunc: func(ctx context.Context, userID uuid.UUID) error {
//				panic("mock out the HardDeleteUser method")
//			},
//			InsertOrUpdateImageFunc: func(ctx context.Context, arg InsertOrUpdateImageParams) error {
//				panic("mock out the InsertOrUpdateImage method")
//			},
//			InsertOrUpdateRolesFunc: func(ctx context.Context, arg InsertOrUpdateRolesParams) error {
//				panic("mock out the InsertOrUpdateRoles method")
//			},
//			InsertOrUpdateUserFunc: func(ctx context.Context, arg InsertOrUpdateUserParams) error {
//				panic("mock out the InsertOrUpdateUser method")
//			},
//			RollbackFunc: func(ctx context.Context) error {
//				panic("mock out the Rollback method")
//			},
//			SoftDeleteUserFunc: func(ctx context.Context, userID uuid.UUID) error {
//				panic("mock out the SoftDeleteUser method")
//			},
//		}
//
//		// use mockedTransactioner in code that requires Transactioner
//		// and then make assertions.
//
//	}
type TransactionerMock struct {
	// CheckImageExistsFunc mocks the CheckImageExists method.
	CheckImageExistsFunc func(ctx context.Context, id uuid.UUID) (bool, error)

	// CheckImageHashExistsFunc mocks the CheckImageHashExists method.
	CheckImageHashExistsFunc func(ctx context.Context, fileHash []byte) (bool, error)

	// CommitFunc mocks the Commit method.
	CommitFunc func(ctx context.Context) error

	// DeleteImageFunc mocks the DeleteImage method.
	DeleteImageFunc func(ctx context.Context, id uuid.UUID) error

	// DeleteRoleFunc mocks the DeleteRole method.
	DeleteRoleFunc func(ctx context.Context, id string) error

	// GetAllImagesInIDSFunc mocks the GetAllImagesInIDS method.
	GetAllImagesInIDSFunc func(ctx context.Context, imageIds []uuid.UUID) ([]GetAllImagesInIDSRow, error)

	// GetAllRolesFunc mocks the GetAllRoles method.
	GetAllRolesFunc func(ctx context.Context) ([]Role, error)

	// GetAllUsersInIDSFunc mocks the GetAllUsersInIDS method.
	GetAllUsersInIDSFunc func(ctx context.Context, userIds []uuid.UUID) ([]GetAllUsersInIDSRow, error)

	// GetImageByIdFunc mocks the GetImageById method.
	GetImageByIdFunc func(ctx context.Context, id uuid.UUID) (Image, error)

	// GetManyImagesFunc mocks the GetManyImages method.
	GetManyImagesFunc func(ctx context.Context, arg GetManyImagesParams) ([]GetManyImagesRow, error)

	// GetManyUsersFunc mocks the GetManyUsers method.
	GetManyUsersFunc func(ctx context.Context, arg GetManyUsersParams) ([]GetManyUsersRow, error)

	// GetUserByEmailFunc mocks the GetUserByEmail method.
	GetUserByEmailFunc func(ctx context.Context, email string) (GetUserByEmailRow, error)

	// GetUserByIdFunc mocks the GetUserById method.
	GetUserByIdFunc func(ctx context.Context, userID uuid.UUID) (GetUserByIdRow, error)

	// GetUsersCountFunc mocks the GetUsersCount method.
	GetUsersCountFunc func(ctx context.Context) (int64, error)

	// GetUsersCountByRoleFunc mocks the GetUsersCountByRole method.
	GetUsersCountByRoleFunc func(ctx context.Context, userRole string) (int64, error)

	// HardDeleteUserFunc mocks the HardDeleteUser method.
	HardDeleteUserFunc func(ctx context.Context, userID uuid.UUID) error

	// InsertOrUpdateImageFunc mocks the InsertOrUpdateImage method.
	InsertOrUpdateImageFunc func(ctx context.Context, arg InsertOrUpdateImageParams) error

	// InsertOrUpdateRolesFunc mocks the InsertOrUpdateRoles method.
	InsertOrUpdateRolesFunc func(ctx context.Context, arg InsertOrUpdateRolesParams) error

	// InsertOrUpdateUserFunc mocks the InsertOrUpdateUser method.
	InsertOrUpdateUserFunc func(ctx context.Context, arg InsertOrUpdateUserParams) error

	// RollbackFunc mocks the Rollback method.
	RollbackFunc func(ctx context.Context) error

	// SoftDeleteUserFunc mocks the SoftDeleteUser method.
	SoftDeleteUserFunc func(ctx context.Context, userID uuid.UUID) error

	// calls tracks calls to the methods.
	calls struct {
		// CheckImageExists holds details about calls to the CheckImageExists method.
		CheckImageExists []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// CheckImageHashExists holds details about calls to the CheckImageHashExists method.
		CheckImageHashExists []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// FileHash is the fileHash argument value.
			FileHash []byte
		}
		// Commit holds details about calls to the Commit method.
		Commit []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// DeleteImage holds details about calls to the DeleteImage method.
		DeleteImage []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// DeleteRole holds details about calls to the DeleteRole method.
		DeleteRole []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// GetAllImagesInIDS holds details about calls to the GetAllImagesInIDS method.
		GetAllImagesInIDS []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ImageIds is the imageIds argument value.
			ImageIds []uuid.UUID
		}
		// GetAllRoles holds details about calls to the GetAllRoles method.
		GetAllRoles []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetAllUsersInIDS holds details about calls to the GetAllUsersInIDS method.
		GetAllUsersInIDS []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserIds is the userIds argument value.
			UserIds []uuid.UUID
		}
		// GetImageById holds details about calls to the GetImageById method.
		GetImageById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// GetManyImages holds details about calls to the GetManyImages method.
		GetManyImages []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg GetManyImagesParams
		}
		// GetManyUsers holds details about calls to the GetManyUsers method.
		GetManyUsers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg GetManyUsersParams
		}
		// GetUserByEmail holds details about calls to the GetUserByEmail method.
		GetUserByEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Email is the email argument value.
			Email string
		}
		// GetUserById holds details about calls to the GetUserById method.
		GetUserById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
		}
		// GetUsersCount holds details about calls to the GetUsersCount method.
		GetUsersCount []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetUsersCountByRole holds details about calls to the GetUsersCountByRole method.
		GetUsersCountByRole []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserRole is the userRole argument value.
			UserRole string
		}
		// HardDeleteUser holds details about calls to the HardDeleteUser method.
		HardDeleteUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
		}
		// InsertOrUpdateImage holds details about calls to the InsertOrUpdateImage method.
		InsertOrUpdateImage []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg InsertOrUpdateImageParams
		}
		// InsertOrUpdateRoles holds details about calls to the InsertOrUpdateRoles method.
		InsertOrUpdateRoles []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg InsertOrUpdateRolesParams
		}
		// InsertOrUpdateUser holds details about calls to the InsertOrUpdateUser method.
		InsertOrUpdateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Arg is the arg argument value.
			Arg InsertOrUpdateUserParams
		}
		// Rollback holds details about calls to the Rollback method.
		Rollback []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// SoftDeleteUser holds details about calls to the SoftDeleteUser method.
		SoftDeleteUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uuid.UUID
		}
	}
	lockCheckImageExists     sync.RWMutex
	lockCheckImageHashExists sync.RWMutex
	lockCommit               sync.RWMutex
	lockDeleteImage          sync.RWMutex
	lockDeleteRole           sync.RWMutex
	lockGetAllImagesInIDS    sync.RWMutex
	lockGetAllRoles          sync.RWMutex
	lockGetAllUsersInIDS     sync.RWMutex
	lockGetImageById         sync.RWMutex
	lockGetManyImages        sync.RWMutex
	lockGetManyUsers         sync.RWMutex
	lockGetUserByEmail       sync.RWMutex
	lockGetUserById          sync.RWMutex
	lockGetUsersCount        sync.RWMutex
	lockGetUsersCountByRole  sync.RWMutex
	lockHardDeleteUser       sync.RWMutex
	lockInsertOrUpdateImage  sync.RWMutex
	lockInsertOrUpdateRoles  sync.RWMutex
	lockInsertOrUpdateUser   sync.RWMutex
	lockRollback             sync.RWMutex
	lockSoftDeleteUser       sync.RWMutex
}

// CheckImageExists calls CheckImageExistsFunc.
func (mock *TransactionerMock) CheckImageExists(ctx context.Context, id uuid.UUID) (bool, error) {
	if mock.CheckImageExistsFunc == nil {
		panic("TransactionerMock.CheckImageExistsFunc: method is nil but Transactioner.CheckImageExists was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockCheckImageExists.Lock()
	mock.calls.CheckImageExists = append(mock.calls.CheckImageExists, callInfo)
	mock.lockCheckImageExists.Unlock()
	return mock.CheckImageExistsFunc(ctx, id)
}

// CheckImageExistsCalls gets all the calls that were made to CheckImageExists.
// Check the length with:
//
//	len(mockedTransactioner.CheckImageExistsCalls())
func (mock *TransactionerMock) CheckImageExistsCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockCheckImageExists.RLock()
	calls = mock.calls.CheckImageExists
	mock.lockCheckImageExists.RUnlock()
	return calls
}

// CheckImageHashExists calls CheckImageHashExistsFunc.
func (mock *TransactionerMock) CheckImageHashExists(ctx context.Context, fileHash []byte) (bool, error) {
	if mock.CheckImageHashExistsFunc == nil {
		panic("TransactionerMock.CheckImageHashExistsFunc: method is nil but Transactioner.CheckImageHashExists was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		FileHash []byte
	}{
		Ctx:      ctx,
		FileHash: fileHash,
	}
	mock.lockCheckImageHashExists.Lock()
	mock.calls.CheckImageHashExists = append(mock.calls.CheckImageHashExists, callInfo)
	mock.lockCheckImageHashExists.Unlock()
	return mock.CheckImageHashExistsFunc(ctx, fileHash)
}

// CheckImageHashExistsCalls gets all the calls that were made to CheckImageHashExists.
// Check the length with:
//
//	len(mockedTransactioner.CheckImageHashExistsCalls())
func (mock *TransactionerMock) CheckImageHashExistsCalls() []struct {
	Ctx      context.Context
	FileHash []byte
} {
	var calls []struct {
		Ctx      context.Context
		FileHash []byte
	}
	mock.lockCheckImageHashExists.RLock()
	calls = mock.calls.CheckImageHashExists
	mock.lockCheckImageHashExists.RUnlock()
	return calls
}

// Commit calls CommitFunc.
func (mock *TransactionerMock) Commit(ctx context.Context) error {
	if mock.CommitFunc == nil {
		panic("TransactionerMock.CommitFunc: method is nil but Transactioner.Commit was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCommit.Lock()
	mock.calls.Commit = append(mock.calls.Commit, callInfo)
	mock.lockCommit.Unlock()
	return mock.CommitFunc(ctx)
}

// CommitCalls gets all the calls that were made to Commit.
// Check the length with:
//
//	len(mockedTransactioner.CommitCalls())
func (mock *TransactionerMock) CommitCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockCommit.RLock()
	calls = mock.calls.Commit
	mock.lockCommit.RUnlock()
	return calls
}

// DeleteImage calls DeleteImageFunc.
func (mock *TransactionerMock) DeleteImage(ctx context.Context, id uuid.UUID) error {
	if mock.DeleteImageFunc == nil {
		panic("TransactionerMock.DeleteImageFunc: method is nil but Transactioner.DeleteImage was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteImage.Lock()
	mock.calls.DeleteImage = append(mock.calls.DeleteImage, callInfo)
	mock.lockDeleteImage.Unlock()
	return mock.DeleteImageFunc(ctx, id)
}

// DeleteImageCalls gets all the calls that were made to DeleteImage.
// Check the length with:
//
//	len(mockedTransactioner.DeleteImageCalls())
func (mock *TransactionerMock) DeleteImageCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockDeleteImage.RLock()
	calls = mock.calls.DeleteImage
	mock.lockDeleteImage.RUnlock()
	return calls
}

// DeleteRole calls DeleteRoleFunc.
func (mock *TransactionerMock) DeleteRole(ctx context.Context, id string) error {
	if mock.DeleteRoleFunc == nil {
		panic("TransactionerMock.DeleteRoleFunc: method is nil but Transactioner.DeleteRole was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteRole.Lock()
	mock.calls.DeleteRole = append(mock.calls.DeleteRole, callInfo)
	mock.lockDeleteRole.Unlock()
	return mock.DeleteRoleFunc(ctx, id)
}

// DeleteRoleCalls gets all the calls that were made to DeleteRole.
// Check the length with:
//
//	len(mockedTransactioner.DeleteRoleCalls())
func (mock *TransactionerMock) DeleteRoleCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockDeleteRole.RLock()
	calls = mock.calls.DeleteRole
	mock.lockDeleteRole.RUnlock()
	return calls
}

// GetAllImagesInIDS calls GetAllImagesInIDSFunc.
func (mock *TransactionerMock) GetAllImagesInIDS(ctx context.Context, imageIds []uuid.UUID) ([]GetAllImagesInIDSRow, error) {
	if mock.GetAllImagesInIDSFunc == nil {
		panic("TransactionerMock.GetAllImagesInIDSFunc: method is nil but Transactioner.GetAllImagesInIDS was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ImageIds []uuid.UUID
	}{
		Ctx:      ctx,
		ImageIds: imageIds,
	}
	mock.lockGetAllImagesInIDS.Lock()
	mock.calls.GetAllImagesInIDS = append(mock.calls.GetAllImagesInIDS, callInfo)
	mock.lockGetAllImagesInIDS.Unlock()
	return mock.GetAllImagesInIDSFunc(ctx, imageIds)
}

// GetAllImagesInIDSCalls gets all the calls that were made to GetAllImagesInIDS.
// Check the length with:
//
//	len(mockedTransactioner.GetAllImagesInIDSCalls())
func (mock *TransactionerMock) GetAllImagesInIDSCalls() []struct {
	Ctx      context.Context
	ImageIds []uuid.UUID
} {
	var calls []struct {
		Ctx      context.Context
		ImageIds []uuid.UUID
	}
	mock.lockGetAllImagesInIDS.RLock()
	calls = mock.calls.GetAllImagesInIDS
	mock.lockGetAllImagesInIDS.RUnlock()
	return calls
}

// GetAllRoles calls GetAllRolesFunc.
func (mock *TransactionerMock) GetAllRoles(ctx context.Context) ([]Role, error) {
	if mock.GetAllRolesFunc == nil {
		panic("TransactionerMock.GetAllRolesFunc: method is nil but Transactioner.GetAllRoles was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAllRoles.Lock()
	mock.calls.GetAllRoles = append(mock.calls.GetAllRoles, callInfo)
	mock.lockGetAllRoles.Unlock()
	return mock.GetAllRolesFunc(ctx)
}

// GetAllRolesCalls gets all the calls that were made to GetAllRoles.
// Check the length with:
//
//	len(mockedTransactioner.GetAllRolesCalls())
func (mock *TransactionerMock) GetAllRolesCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAllRoles.RLock()
	calls = mock.calls.GetAllRoles
	mock.lockGetAllRoles.RUnlock()
	return calls
}

// GetAllUsersInIDS calls GetAllUsersInIDSFunc.
func (mock *TransactionerMock) GetAllUsersInIDS(ctx context.Context, userIds []uuid.UUID) ([]GetAllUsersInIDSRow, error) {
	if mock.GetAllUsersInIDSFunc == nil {
		panic("TransactionerMock.GetAllUsersInIDSFunc: method is nil but Transactioner.GetAllUsersInIDS was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		UserIds []uuid.UUID
	}{
		Ctx:     ctx,
		UserIds: userIds,
	}
	mock.lockGetAllUsersInIDS.Lock()
	mock.calls.GetAllUsersInIDS = append(mock.calls.GetAllUsersInIDS, callInfo)
	mock.lockGetAllUsersInIDS.Unlock()
	return mock.GetAllUsersInIDSFunc(ctx, userIds)
}

// GetAllUsersInIDSCalls gets all the calls that were made to GetAllUsersInIDS.
// Check the length with:
//
//	len(mockedTransactioner.GetAllUsersInIDSCalls())
func (mock *TransactionerMock) GetAllUsersInIDSCalls() []struct {
	Ctx     context.Context
	UserIds []uuid.UUID
} {
	var calls []struct {
		Ctx     context.Context
		UserIds []uuid.UUID
	}
	mock.lockGetAllUsersInIDS.RLock()
	calls = mock.calls.GetAllUsersInIDS
	mock.lockGetAllUsersInIDS.RUnlock()
	return calls
}

// GetImageById calls GetImageByIdFunc.
func (mock *TransactionerMock) GetImageById(ctx context.Context, id uuid.UUID) (Image, error) {
	if mock.GetImageByIdFunc == nil {
		panic("TransactionerMock.GetImageByIdFunc: method is nil but Transactioner.GetImageById was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetImageById.Lock()
	mock.calls.GetImageById = append(mock.calls.GetImageById, callInfo)
	mock.lockGetImageById.Unlock()
	return mock.GetImageByIdFunc(ctx, id)
}

// GetImageByIdCalls gets all the calls that were made to GetImageById.
// Check the length with:
//
//	len(mockedTransactioner.GetImageByIdCalls())
func (mock *TransactionerMock) GetImageByIdCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockGetImageById.RLock()
	calls = mock.calls.GetImageById
	mock.lockGetImageById.RUnlock()
	return calls
}

// GetManyImages calls GetManyImagesFunc.
func (mock *TransactionerMock) GetManyImages(ctx context.Context, arg GetManyImagesParams) ([]GetManyImagesRow, error) {
	if mock.GetManyImagesFunc == nil {
		panic("TransactionerMock.GetManyImagesFunc: method is nil but Transactioner.GetManyImages was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg GetManyImagesParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockGetManyImages.Lock()
	mock.calls.GetManyImages = append(mock.calls.GetManyImages, callInfo)
	mock.lockGetManyImages.Unlock()
	return mock.GetManyImagesFunc(ctx, arg)
}

// GetManyImagesCalls gets all the calls that were made to GetManyImages.
// Check the length with:
//
//	len(mockedTransactioner.GetManyImagesCalls())
func (mock *TransactionerMock) GetManyImagesCalls() []struct {
	Ctx context.Context
	Arg GetManyImagesParams
} {
	var calls []struct {
		Ctx context.Context
		Arg GetManyImagesParams
	}
	mock.lockGetManyImages.RLock()
	calls = mock.calls.GetManyImages
	mock.lockGetManyImages.RUnlock()
	return calls
}

// GetManyUsers calls GetManyUsersFunc.
func (mock *TransactionerMock) GetManyUsers(ctx context.Context, arg GetManyUsersParams) ([]GetManyUsersRow, error) {
	if mock.GetManyUsersFunc == nil {
		panic("TransactionerMock.GetManyUsersFunc: method is nil but Transactioner.GetManyUsers was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg GetManyUsersParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockGetManyUsers.Lock()
	mock.calls.GetManyUsers = append(mock.calls.GetManyUsers, callInfo)
	mock.lockGetManyUsers.Unlock()
	return mock.GetManyUsersFunc(ctx, arg)
}

// GetManyUsersCalls gets all the calls that were made to GetManyUsers.
// Check the length with:
//
//	len(mockedTransactioner.GetManyUsersCalls())
func (mock *TransactionerMock) GetManyUsersCalls() []struct {
	Ctx context.Context
	Arg GetManyUsersParams
} {
	var calls []struct {
		Ctx context.Context
		Arg GetManyUsersParams
	}
	mock.lockGetManyUsers.RLock()
	calls = mock.calls.GetManyUsers
	mock.lockGetManyUsers.RUnlock()
	return calls
}

// GetUserByEmail calls GetUserByEmailFunc.
func (mock *TransactionerMock) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	if mock.GetUserByEmailFunc == nil {
		panic("TransactionerMock.GetUserByEmailFunc: method is nil but Transactioner.GetUserByEmail was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Email string
	}{
		Ctx:   ctx,
		Email: email,
	}
	mock.lockGetUserByEmail.Lock()
	mock.calls.GetUserByEmail = append(mock.calls.GetUserByEmail, callInfo)
	mock.lockGetUserByEmail.Unlock()
	return mock.GetUserByEmailFunc(ctx, email)
}

// GetUserByEmailCalls gets all the calls that were made to GetUserByEmail.
// Check the length with:
//
//	len(mockedTransactioner.GetUserByEmailCalls())
func (mock *TransactionerMock) GetUserByEmailCalls() []struct {
	Ctx   context.Context
	Email string
} {
	var calls []struct {
		Ctx   context.Context
		Email string
	}
	mock.lockGetUserByEmail.RLock()
	calls = mock.calls.GetUserByEmail
	mock.lockGetUserByEmail.RUnlock()
	return calls
}

// GetUserById calls GetUserByIdFunc.
func (mock *TransactionerMock) GetUserById(ctx context.Context, userID uuid.UUID) (GetUserByIdRow, error) {
	if mock.GetUserByIdFunc == nil {
		panic("TransactionerMock.GetUserByIdFunc: method is nil but Transactioner.GetUserById was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockGetUserById.Lock()
	mock.calls.GetUserById = append(mock.calls.GetUserById, callInfo)
	mock.lockGetUserById.Unlock()
	return mock.GetUserByIdFunc(ctx, userID)
}

// GetUserByIdCalls gets all the calls that were made to GetUserById.
// Check the length with:
//
//	len(mockedTransactioner.GetUserByIdCalls())
func (mock *TransactionerMock) GetUserByIdCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
	}
	mock.lockGetUserById.RLock()
	calls = mock.calls.GetUserById
	mock.lockGetUserById.RUnlock()
	return calls
}

// GetUsersCount calls GetUsersCountFunc.
func (mock *TransactionerMock) GetUsersCount(ctx context.Context) (int64, error) {
	if mock.GetUsersCountFunc == nil {
		panic("TransactionerMock.GetUsersCountFunc: method is nil but Transactioner.GetUsersCount was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetUsersCount.Lock()
	mock.calls.GetUsersCount = append(mock.calls.GetUsersCount, callInfo)
	mock.lockGetUsersCount.Unlock()
	return mock.GetUsersCountFunc(ctx)
}

// GetUsersCountCalls gets all the calls that were made to GetUsersCount.
// Check the length with:
//
//	len(mockedTransactioner.GetUsersCountCalls())
func (mock *TransactionerMock) GetUsersCountCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetUsersCount.RLock()
	calls = mock.calls.GetUsersCount
	mock.lockGetUsersCount.RUnlock()
	return calls
}

// GetUsersCountByRole calls GetUsersCountByRoleFunc.
func (mock *TransactionerMock) GetUsersCountByRole(ctx context.Context, userRole string) (int64, error) {
	if mock.GetUsersCountByRoleFunc == nil {
		panic("TransactionerMock.GetUsersCountByRoleFunc: method is nil but Transactioner.GetUsersCountByRole was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		UserRole string
	}{
		Ctx:      ctx,
		UserRole: userRole,
	}
	mock.lockGetUsersCountByRole.Lock()
	mock.calls.GetUsersCountByRole = append(mock.calls.GetUsersCountByRole, callInfo)
	mock.lockGetUsersCountByRole.Unlock()
	return mock.GetUsersCountByRoleFunc(ctx, userRole)
}

// GetUsersCountByRoleCalls gets all the calls that were made to GetUsersCountByRole.
// Check the length with:
//
//	len(mockedTransactioner.GetUsersCountByRoleCalls())
func (mock *TransactionerMock) GetUsersCountByRoleCalls() []struct {
	Ctx      context.Context
	UserRole string
} {
	var calls []struct {
		Ctx      context.Context
		UserRole string
	}
	mock.lockGetUsersCountByRole.RLock()
	calls = mock.calls.GetUsersCountByRole
	mock.lockGetUsersCountByRole.RUnlock()
	return calls
}

// HardDeleteUser calls HardDeleteUserFunc.
func (mock *TransactionerMock) HardDeleteUser(ctx context.Context, userID uuid.UUID) error {
	if mock.HardDeleteUserFunc == nil {
		panic("TransactionerMock.HardDeleteUserFunc: method is nil but Transactioner.HardDeleteUser was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockHardDeleteUser.Lock()
	mock.calls.HardDeleteUser = append(mock.calls.HardDeleteUser, callInfo)
	mock.lockHardDeleteUser.Unlock()
	return mock.HardDeleteUserFunc(ctx, userID)
}

// HardDeleteUserCalls gets all the calls that were made to HardDeleteUser.
// Check the length with:
//
//	len(mockedTransactioner.HardDeleteUserCalls())
func (mock *TransactionerMock) HardDeleteUserCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
	}
	mock.lockHardDeleteUser.RLock()
	calls = mock.calls.HardDeleteUser
	mock.lockHardDeleteUser.RUnlock()
	return calls
}

// InsertOrUpdateImage calls InsertOrUpdateImageFunc.
func (mock *TransactionerMock) InsertOrUpdateImage(ctx context.Context, arg InsertOrUpdateImageParams) error {
	if mock.InsertOrUpdateImageFunc == nil {
		panic("TransactionerMock.InsertOrUpdateImageFunc: method is nil but Transactioner.InsertOrUpdateImage was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg InsertOrUpdateImageParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockInsertOrUpdateImage.Lock()
	mock.calls.InsertOrUpdateImage = append(mock.calls.InsertOrUpdateImage, callInfo)
	mock.lockInsertOrUpdateImage.Unlock()
	return mock.InsertOrUpdateImageFunc(ctx, arg)
}

// InsertOrUpdateImageCalls gets all the calls that were made to InsertOrUpdateImage.
// Check the length with:
//
//	len(mockedTransactioner.InsertOrUpdateImageCalls())
func (mock *TransactionerMock) InsertOrUpdateImageCalls() []struct {
	Ctx context.Context
	Arg InsertOrUpdateImageParams
} {
	var calls []struct {
		Ctx context.Context
		Arg InsertOrUpdateImageParams
	}
	mock.lockInsertOrUpdateImage.RLock()
	calls = mock.calls.InsertOrUpdateImage
	mock.lockInsertOrUpdateImage.RUnlock()
	return calls
}

// InsertOrUpdateRoles calls InsertOrUpdateRolesFunc.
func (mock *TransactionerMock) InsertOrUpdateRoles(ctx context.Context, arg InsertOrUpdateRolesParams) error {
	if mock.InsertOrUpdateRolesFunc == nil {
		panic("TransactionerMock.InsertOrUpdateRolesFunc: method is nil but Transactioner.InsertOrUpdateRoles was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg InsertOrUpdateRolesParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockInsertOrUpdateRoles.Lock()
	mock.calls.InsertOrUpdateRoles = append(mock.calls.InsertOrUpdateRoles, callInfo)
	mock.lockInsertOrUpdateRoles.Unlock()
	return mock.InsertOrUpdateRolesFunc(ctx, arg)
}

// InsertOrUpdateRolesCalls gets all the calls that were made to InsertOrUpdateRoles.
// Check the length with:
//
//	len(mockedTransactioner.InsertOrUpdateRolesCalls())
func (mock *TransactionerMock) InsertOrUpdateRolesCalls() []struct {
	Ctx context.Context
	Arg InsertOrUpdateRolesParams
} {
	var calls []struct {
		Ctx context.Context
		Arg InsertOrUpdateRolesParams
	}
	mock.lockInsertOrUpdateRoles.RLock()
	calls = mock.calls.InsertOrUpdateRoles
	mock.lockInsertOrUpdateRoles.RUnlock()
	return calls
}

// InsertOrUpdateUser calls InsertOrUpdateUserFunc.
func (mock *TransactionerMock) InsertOrUpdateUser(ctx context.Context, arg InsertOrUpdateUserParams) error {
	if mock.InsertOrUpdateUserFunc == nil {
		panic("TransactionerMock.InsertOrUpdateUserFunc: method is nil but Transactioner.InsertOrUpdateUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Arg InsertOrUpdateUserParams
	}{
		Ctx: ctx,
		Arg: arg,
	}
	mock.lockInsertOrUpdateUser.Lock()
	mock.calls.InsertOrUpdateUser = append(mock.calls.InsertOrUpdateUser, callInfo)
	mock.lockInsertOrUpdateUser.Unlock()
	return mock.InsertOrUpdateUserFunc(ctx, arg)
}

// InsertOrUpdateUserCalls gets all the calls that were made to InsertOrUpdateUser.
// Check the length with:
//
//	len(mockedTransactioner.InsertOrUpdateUserCalls())
func (mock *TransactionerMock) InsertOrUpdateUserCalls() []struct {
	Ctx context.Context
	Arg InsertOrUpdateUserParams
} {
	var calls []struct {
		Ctx context.Context
		Arg InsertOrUpdateUserParams
	}
	mock.lockInsertOrUpdateUser.RLock()
	calls = mock.calls.InsertOrUpdateUser
	mock.lockInsertOrUpdateUser.RUnlock()
	return calls
}

// Rollback calls RollbackFunc.
func (mock *TransactionerMock) Rollback(ctx context.Context) error {
	if mock.RollbackFunc == nil {
		panic("TransactionerMock.RollbackFunc: method is nil but Transactioner.Rollback was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockRollback.Lock()
	mock.calls.Rollback = append(mock.calls.Rollback, callInfo)
	mock.lockRollback.Unlock()
	return mock.RollbackFunc(ctx)
}

// RollbackCalls gets all the calls that were made to Rollback.
// Check the length with:
//
//	len(mockedTransactioner.RollbackCalls())
func (mock *TransactionerMock) RollbackCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockRollback.RLock()
	calls = mock.calls.Rollback
	mock.lockRollback.RUnlock()
	return calls
}

// SoftDeleteUser calls SoftDeleteUserFunc.
func (mock *TransactionerMock) SoftDeleteUser(ctx context.Context, userID uuid.UUID) error {
	if mock.SoftDeleteUserFunc == nil {
		panic("TransactionerMock.SoftDeleteUserFunc: method is nil but Transactioner.SoftDeleteUser was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uuid.UUID
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	mock.lockSoftDeleteUser.Lock()
	mock.calls.SoftDeleteUser = append(mock.calls.SoftDeleteUser, callInfo)
	mock.lockSoftDeleteUser.Unlock()
	return mock.SoftDeleteUserFunc(ctx, userID)
}

// SoftDeleteUserCalls gets all the calls that were made to SoftDeleteUser.
// Check the length with:
//
//	len(mockedTransactioner.SoftDeleteUserCalls())
func (mock *TransactionerMock) SoftDeleteUserCalls() []struct {
	Ctx    context.Context
	UserID uuid.UUID
} {
	var calls []struct {
		Ctx    context.Context
		UserID uuid.UUID
	}
	mock.lockSoftDeleteUser.RLock()
	calls = mock.calls.SoftDeleteUser
	mock.lockSoftDeleteUser.RUnlock()
	return calls
}