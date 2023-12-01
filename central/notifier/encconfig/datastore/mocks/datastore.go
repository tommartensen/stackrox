// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/datastore.go -source datastore.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockDataStore) GetConfig() (*storage.NotifierEncConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*storage.NotifierEncConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockDataStoreMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockDataStore)(nil).GetConfig))
}

// UpsertConfig mocks base method.
func (m *MockDataStore) UpsertConfig(config *storage.NotifierEncConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertConfig", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertConfig indicates an expected call of UpsertConfig.
func (mr *MockDataStoreMockRecorder) UpsertConfig(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertConfig", reflect.TypeOf((*MockDataStore)(nil).UpsertConfig), config)
}