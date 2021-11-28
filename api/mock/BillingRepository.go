// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/repository/billingRepository/billingRepository.go

// Package mock is a generated GoMock package.
package mock

import (
	billingModel "goa-golang/app/model/billingModel"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBillingRepositoryInterface is a mock of BillingRepositoryInterface interface.
type MockBillingRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBillingRepositoryInterfaceMockRecorder
}

// MockBillingRepositoryInterfaceMockRecorder is the mock recorder for MockBillingRepositoryInterface.
type MockBillingRepositoryInterfaceMockRecorder struct {
	mock *MockBillingRepositoryInterface
}

// NewMockBillingRepositoryInterface creates a new mock instance.
func NewMockBillingRepositoryInterface(ctrl *gomock.Controller) *MockBillingRepositoryInterface {
	mock := &MockBillingRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockBillingRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBillingRepositoryInterface) EXPECT() *MockBillingRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateBillingService mocks base method.
func (m *MockBillingRepositoryInterface) CreateBillingService(identity billingModel.Identify, key, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBillingService", identity, key, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBillingService indicates an expected call of CreateBillingService.
func (mr *MockBillingRepositoryInterfaceMockRecorder) CreateBillingService(identity, key, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBillingService", reflect.TypeOf((*MockBillingRepositoryInterface)(nil).CreateBillingService), identity, key, userID)
}
