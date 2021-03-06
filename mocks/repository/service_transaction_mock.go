// Code generated by MockGen. DO NOT EDIT.
// Source: ./repositories/service_transaction_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "AntarJemput-Be-C/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceTransacInterface is a mock of ServiceTransacInterface interface.
type MockServiceTransacInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceTransacInterfaceMockRecorder
}

// MockServiceTransacInterfaceMockRecorder is the mock recorder for MockServiceTransacInterface.
type MockServiceTransacInterfaceMockRecorder struct {
	mock *MockServiceTransacInterface
}

// NewMockServiceTransacInterface creates a new mock instance.
func NewMockServiceTransacInterface(ctrl *gomock.Controller) *MockServiceTransacInterface {
	mock := &MockServiceTransacInterface{ctrl: ctrl}
	mock.recorder = &MockServiceTransacInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceTransacInterface) EXPECT() *MockServiceTransacInterfaceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockServiceTransacInterface) GetAll() ([]models.ServiceTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.ServiceTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockServiceTransacInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockServiceTransacInterface)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockServiceTransacInterface) GetById(id int) (models.ServiceTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(models.ServiceTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockServiceTransacInterfaceMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockServiceTransacInterface)(nil).GetById), id)
}

// Save mocks base method.
func (m *MockServiceTransacInterface) Save(serviceTransaction *models.ServiceTransaction) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", serviceTransaction)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockServiceTransacInterfaceMockRecorder) Save(serviceTransaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockServiceTransacInterface)(nil).Save), serviceTransaction)
}
