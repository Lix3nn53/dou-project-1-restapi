// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/service/userService.go

// Package mock is a generated GoMock package.
package mock

import (
	"dou-survey/app/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserServiceInterface is a mock of UserServiceInterface interface.
type MockUserServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceInterfaceMockRecorder
}

// MockUserServiceInterfaceMockRecorder is the mock recorder for MockUserServiceInterface.
type MockUserServiceInterfaceMockRecorder struct {
	mock *MockUserServiceInterface
}

// NewMockUserServiceInterface creates a new mock instance.
func NewMockUserServiceInterface(ctrl *gomock.Controller) *MockUserServiceInterface {
	mock := &MockUserServiceInterface{ctrl: ctrl}
	mock.recorder = &MockUserServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceInterface) EXPECT() *MockUserServiceInterfaceMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockUserServiceInterface) FindByID(id, field string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id, field)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserServiceInterfaceMockRecorder) FindByID(id, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserServiceInterface)(nil).FindByID), id, field)
}
