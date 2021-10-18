// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/serviceTransaction_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "AntarJemput-Be-C/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSTSInterface is a mock of STSInterface interface.
type MockSTSInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSTSInterfaceMockRecorder
}

// MockSTSInterfaceMockRecorder is the mock recorder for MockSTSInterface.
type MockSTSInterfaceMockRecorder struct {
	mock *MockSTSInterface
}

// NewMockSTSInterface creates a new mock instance.
func NewMockSTSInterface(ctrl *gomock.Controller) *MockSTSInterface {
	mock := &MockSTSInterface{ctrl: ctrl}
	mock.recorder = &MockSTSInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSTSInterface) EXPECT() *MockSTSInterfaceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockSTSInterface) GetAll() ([]models.ServiceTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.ServiceTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockSTSInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockSTSInterface)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockSTSInterface) GetById(id int) (models.ServiceTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(models.ServiceTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockSTSInterfaceMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockSTSInterface)(nil).GetById), id)
}

// Save mocks base method.
func (m *MockSTSInterface) Save(st *models.ServiceTransaction) (*models.ServiceTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", st)
	ret0, _ := ret[0].(*models.ServiceTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockSTSInterfaceMockRecorder) Save(st interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockSTSInterface)(nil).Save), st)
}
