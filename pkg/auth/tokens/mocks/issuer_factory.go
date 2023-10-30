// Code generated by MockGen. DO NOT EDIT.
// Source: issuer_factory.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/issuer_factory.go -source issuer_factory.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	tokens "github.com/stackrox/rox/pkg/auth/tokens"
	gomock "go.uber.org/mock/gomock"
)

// MockIssuerFactory is a mock of IssuerFactory interface.
type MockIssuerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockIssuerFactoryMockRecorder
}

// MockIssuerFactoryMockRecorder is the mock recorder for MockIssuerFactory.
type MockIssuerFactoryMockRecorder struct {
	mock *MockIssuerFactory
}

// NewMockIssuerFactory creates a new mock instance.
func NewMockIssuerFactory(ctrl *gomock.Controller) *MockIssuerFactory {
	mock := &MockIssuerFactory{ctrl: ctrl}
	mock.recorder = &MockIssuerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssuerFactory) EXPECT() *MockIssuerFactoryMockRecorder {
	return m.recorder
}

// CreateIssuer mocks base method.
func (m *MockIssuerFactory) CreateIssuer(source tokens.Source, options ...tokens.Option) (tokens.Issuer, error) {
	m.ctrl.T.Helper()
	varargs := []any{source}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateIssuer", varargs...)
	ret0, _ := ret[0].(tokens.Issuer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIssuer indicates an expected call of CreateIssuer.
func (mr *MockIssuerFactoryMockRecorder) CreateIssuer(source any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{source}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIssuer", reflect.TypeOf((*MockIssuerFactory)(nil).CreateIssuer), varargs...)
}

// UnregisterSource mocks base method.
func (m *MockIssuerFactory) UnregisterSource(source tokens.Source) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterSource", source)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterSource indicates an expected call of UnregisterSource.
func (mr *MockIssuerFactoryMockRecorder) UnregisterSource(source any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterSource", reflect.TypeOf((*MockIssuerFactory)(nil).UnregisterSource), source)
}