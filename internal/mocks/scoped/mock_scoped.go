// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fluffy-bunny/grpcdotnetgo-example/internal/contracts/scoped (interfaces: IScoped)

// Package scoped is a generated GoMock package.
package scoped

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIScoped is a mock of IScoped interface.
type MockIScoped struct {
	ctrl     *gomock.Controller
	recorder *MockIScopedMockRecorder
}

// MockIScopedMockRecorder is the mock recorder for MockIScoped.
type MockIScopedMockRecorder struct {
	mock *MockIScoped
}

// NewMockIScoped creates a new mock instance.
func NewMockIScoped(ctrl *gomock.Controller) *MockIScoped {
	mock := &MockIScoped{ctrl: ctrl}
	mock.recorder = &MockIScopedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIScoped) EXPECT() *MockIScopedMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockIScoped) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockIScopedMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIScoped)(nil).GetName))
}

// SetName mocks base method.
func (m *MockIScoped) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MockIScopedMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockIScoped)(nil).SetName), arg0)
}
