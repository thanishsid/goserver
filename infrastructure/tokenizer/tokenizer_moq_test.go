// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package tokenizer

import (
	"context"
	"sync"
)

// Ensure, that TokenizerMock does implement Tokenizer.
// If this is not the case, regenerate this file with moq.
var _ Tokenizer = &TokenizerMock{}

// TokenizerMock is a mock implementation of Tokenizer.
//
//	func TestSomethingThatUsesTokenizer(t *testing.T) {
//
//		// make and configure a mocked Tokenizer
//		mockedTokenizer := &TokenizerMock{
//			CreateTokenFunc: func(ctx context.Context, claims Validateable) (string, error) {
//				panic("mock out the CreateToken method")
//			},
//			GetClaimsFunc: func(ctx context.Context, secureToken string, claims Validateable) error {
//				panic("mock out the GetClaims method")
//			},
//			GetClaimsUnsafeFunc: func(ctx context.Context, secureToken string, claims Validateable) error {
//				panic("mock out the GetClaimsUnsafe method")
//			},
//		}
//
//		// use mockedTokenizer in code that requires Tokenizer
//		// and then make assertions.
//
//	}
type TokenizerMock struct {
	// CreateTokenFunc mocks the CreateToken method.
	CreateTokenFunc func(ctx context.Context, claims Validateable) (string, error)

	// GetClaimsFunc mocks the GetClaims method.
	GetClaimsFunc func(ctx context.Context, secureToken string, claims Validateable) error

	// GetClaimsUnsafeFunc mocks the GetClaimsUnsafe method.
	GetClaimsUnsafeFunc func(ctx context.Context, secureToken string, claims Validateable) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateToken holds details about calls to the CreateToken method.
		CreateToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Claims is the claims argument value.
			Claims Validateable
		}
		// GetClaims holds details about calls to the GetClaims method.
		GetClaims []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SecureToken is the secureToken argument value.
			SecureToken string
			// Claims is the claims argument value.
			Claims Validateable
		}
		// GetClaimsUnsafe holds details about calls to the GetClaimsUnsafe method.
		GetClaimsUnsafe []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SecureToken is the secureToken argument value.
			SecureToken string
			// Claims is the claims argument value.
			Claims Validateable
		}
	}
	lockCreateToken     sync.RWMutex
	lockGetClaims       sync.RWMutex
	lockGetClaimsUnsafe sync.RWMutex
}

// CreateToken calls CreateTokenFunc.
func (mock *TokenizerMock) CreateToken(ctx context.Context, claims Validateable) (string, error) {
	if mock.CreateTokenFunc == nil {
		panic("TokenizerMock.CreateTokenFunc: method is nil but Tokenizer.CreateToken was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Claims Validateable
	}{
		Ctx:    ctx,
		Claims: claims,
	}
	mock.lockCreateToken.Lock()
	mock.calls.CreateToken = append(mock.calls.CreateToken, callInfo)
	mock.lockCreateToken.Unlock()
	return mock.CreateTokenFunc(ctx, claims)
}

// CreateTokenCalls gets all the calls that were made to CreateToken.
// Check the length with:
//
//	len(mockedTokenizer.CreateTokenCalls())
func (mock *TokenizerMock) CreateTokenCalls() []struct {
	Ctx    context.Context
	Claims Validateable
} {
	var calls []struct {
		Ctx    context.Context
		Claims Validateable
	}
	mock.lockCreateToken.RLock()
	calls = mock.calls.CreateToken
	mock.lockCreateToken.RUnlock()
	return calls
}

// GetClaims calls GetClaimsFunc.
func (mock *TokenizerMock) GetClaims(ctx context.Context, secureToken string, claims Validateable) error {
	if mock.GetClaimsFunc == nil {
		panic("TokenizerMock.GetClaimsFunc: method is nil but Tokenizer.GetClaims was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		SecureToken string
		Claims      Validateable
	}{
		Ctx:         ctx,
		SecureToken: secureToken,
		Claims:      claims,
	}
	mock.lockGetClaims.Lock()
	mock.calls.GetClaims = append(mock.calls.GetClaims, callInfo)
	mock.lockGetClaims.Unlock()
	return mock.GetClaimsFunc(ctx, secureToken, claims)
}

// GetClaimsCalls gets all the calls that were made to GetClaims.
// Check the length with:
//
//	len(mockedTokenizer.GetClaimsCalls())
func (mock *TokenizerMock) GetClaimsCalls() []struct {
	Ctx         context.Context
	SecureToken string
	Claims      Validateable
} {
	var calls []struct {
		Ctx         context.Context
		SecureToken string
		Claims      Validateable
	}
	mock.lockGetClaims.RLock()
	calls = mock.calls.GetClaims
	mock.lockGetClaims.RUnlock()
	return calls
}

// GetClaimsUnsafe calls GetClaimsUnsafeFunc.
func (mock *TokenizerMock) GetClaimsUnsafe(ctx context.Context, secureToken string, claims Validateable) error {
	if mock.GetClaimsUnsafeFunc == nil {
		panic("TokenizerMock.GetClaimsUnsafeFunc: method is nil but Tokenizer.GetClaimsUnsafe was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		SecureToken string
		Claims      Validateable
	}{
		Ctx:         ctx,
		SecureToken: secureToken,
		Claims:      claims,
	}
	mock.lockGetClaimsUnsafe.Lock()
	mock.calls.GetClaimsUnsafe = append(mock.calls.GetClaimsUnsafe, callInfo)
	mock.lockGetClaimsUnsafe.Unlock()
	return mock.GetClaimsUnsafeFunc(ctx, secureToken, claims)
}

// GetClaimsUnsafeCalls gets all the calls that were made to GetClaimsUnsafe.
// Check the length with:
//
//	len(mockedTokenizer.GetClaimsUnsafeCalls())
func (mock *TokenizerMock) GetClaimsUnsafeCalls() []struct {
	Ctx         context.Context
	SecureToken string
	Claims      Validateable
} {
	var calls []struct {
		Ctx         context.Context
		SecureToken string
		Claims      Validateable
	}
	mock.lockGetClaimsUnsafe.RLock()
	calls = mock.calls.GetClaimsUnsafe
	mock.lockGetClaimsUnsafe.RUnlock()
	return calls
}