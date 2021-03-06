// Code generated by MockGen. DO NOT EDIT.
// Source: default.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDefault is a mock of Default interface
type MockDefault struct {
	ctrl     *gomock.Controller
	recorder *MockDefaultMockRecorder
}

// MockDefaultMockRecorder is the mock recorder for MockDefault
type MockDefaultMockRecorder struct {
	mock *MockDefault
}

// NewMockDefault creates a new mock instance
func NewMockDefault(ctrl *gomock.Controller) *MockDefault {
	mock := &MockDefault{ctrl: ctrl}
	mock.recorder = &MockDefaultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDefault) EXPECT() *MockDefaultMockRecorder {
	return m.recorder
}

// Ping mocks base method
func (m *MockDefault) Ping() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ping")
}

// Ping indicates an expected call of Ping
func (mr *MockDefaultMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDefault)(nil).Ping))
}
