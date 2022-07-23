// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/thanishsid/goserver/infrastructure/mailer (interfaces: Mailer)

// Package mockmailer is a generated GoMock package.
package mockmailer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mailer "github.com/thanishsid/goserver/infrastructure/mailer"
)

// MockMailer is a mock of Mailer interface.
type MockMailer struct {
	ctrl     *gomock.Controller
	recorder *MockMailerMockRecorder
}

// MockMailerMockRecorder is the mock recorder for MockMailer.
type MockMailerMockRecorder struct {
	mock *MockMailer
}

// NewMockMailer creates a new mock instance.
func NewMockMailer(ctrl *gomock.Controller) *MockMailer {
	mock := &MockMailer{ctrl: ctrl}
	mock.recorder = &MockMailerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailer) EXPECT() *MockMailerMockRecorder {
	return m.recorder
}

// SendLinkMail mocks base method.
func (m *MockMailer) SendLinkMail(arg0 context.Context, arg1 mailer.LinkMailTemplateData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendLinkMail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendLinkMail indicates an expected call of SendLinkMail.
func (mr *MockMailerMockRecorder) SendLinkMail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendLinkMail", reflect.TypeOf((*MockMailer)(nil).SendLinkMail), arg0, arg1)
}
