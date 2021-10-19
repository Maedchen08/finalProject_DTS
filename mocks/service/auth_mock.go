// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/authservice.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "AntarJemput-Be-C/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthServiceInterface is a mock of AuthServiceInterface interface.
type MockAuthServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceInterfaceMockRecorder
}

// MockAuthServiceInterfaceMockRecorder is the mock recorder for MockAuthServiceInterface.
type MockAuthServiceInterfaceMockRecorder struct {
	mock *MockAuthServiceInterface
}

// NewMockAuthServiceInterface creates a new mock instance.
func NewMockAuthServiceInterface(ctrl *gomock.Controller) *MockAuthServiceInterface {
	mock := &MockAuthServiceInterface{ctrl: ctrl}
	mock.recorder = &MockAuthServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceInterface) EXPECT() *MockAuthServiceInterfaceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockAuthServiceInterface) GetAll() ([]models.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAuthServiceInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAuthServiceInterface)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockAuthServiceInterface) GetById(id int) (models.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockAuthServiceInterfaceMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockAuthServiceInterface)(nil).GetById), id)
}

// Login mocks base method.
func (m *MockAuthServiceInterface) Login(username, password string) (models.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", username, password)
	ret0, _ := ret[0].(models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceInterfaceMockRecorder) Login(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthServiceInterface)(nil).Login), username, password)
}

// Register mocks base method.
func (m *MockAuthServiceInterface) Register(users *models.Users) (*models.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", users)
	ret0, _ := ret[0].(*models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthServiceInterfaceMockRecorder) Register(users interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthServiceInterface)(nil).Register), users)
}
