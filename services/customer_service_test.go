package services

import (
	mocksCustomer "AntarJemput-Be-C/mocks/repository"
	"AntarJemput-Be-C/models"
	"errors"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
)

var (
	mockCustomerCtrl        *gomock.Controller
	repoCustomerMock        *mocksCustomer.MockCustomerRepoInterface
	serviceCustomer         *CustomerService
	defaultCustomerInt      = 0
	defaultCustomerModel    models.Customer
	defaultAllCustomerModel []models.Customer
	errCustomerInternal     = errors.New("unexpected system error")
)

func InitMockCustomerRepo(t *testing.T) {
	mockCustomerCtrl = gomock.NewController(t)
	repoCustomerMock = mocksCustomer.NewMockCustomerRepoInterface(mockCustomerCtrl)
	serviceCustomer = NewCustomerService(repoCustomerMock)
}

func TestCustomerService_Save(t *testing.T) {
	InitMockCustomerRepo(t)
	defer mockCustomerCtrl.Finish()

	expectedId := 1
	payload := &models.Customer{
		Id:   1,
		Name: "Atmim",
		NoWa: "087692716291",
	}
	expectedResponse := payload
	expectedResponse.Id = expectedId

	t.Run("should return success", func(t *testing.T) {
		repoCustomerMock.EXPECT().Save(gomock.Any()).Return(expectedId, nil)
		actualResponse, err := serviceCustomer.Save(payload)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("should return error", func(t *testing.T) {
		repoCustomerMock.EXPECT().Save(gomock.Any()).Return(defaultCustomerInt, errCustomerInternal)
		response, actualErr := serviceCustomer.Save(payload)
		assert.Nil(t, response)
		assert.Equal(t, actualErr, errCustomerInternal)
	})
}

func TestCustomerService_GetById(t *testing.T) {
	InitMockCustomerRepo(t)
	defer mockCustomerCtrl.Finish()

	inputId := 1
	payload := models.Customer{
		Id:   1,
		Name: "Atmim",
		NoWa: "087692716291",
	}

	t.Run("should return success", func(t *testing.T) {
		repoCustomerMock.EXPECT().GetById(gomock.Any()).Return(payload, nil)
		actualResponse, err := serviceCustomer.GetById(inputId)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, payload)

	})

	t.Run("should return error", func(t *testing.T) {
		repoCustomerMock.EXPECT().GetById(gomock.Any()).Return(defaultCustomerModel, errCustomerInternal)
		response, actualErr := serviceCustomer.GetById(inputId)
		assert.Equal(t, response, defaultCustomerModel)
		assert.Equal(t, actualErr, errCustomerInternal)
	})

}

func TestCustomerService_GetAll(t *testing.T) {
	InitMockCustomerRepo(t)
	defer mockCustomerCtrl.Finish()

	payload := []models.Customer{
		{
			Id  :   1,
			Name: "Atmim",
			NoWa: "087692716291",
		},
		{
			Id  :   2,
			Name: "Arin",
			NoWa: "087692716291",
		},
		{
			Id  :   3,
			Name: "Tuti Wulandari",
			NoWa: "087692716291",
		},
	}

	t.Run("should return success",func(t *testing.T){
		repoCustomerMock.EXPECT().GetAll().Return(payload, nil)
		actualResponse, err := serviceCustomer.GetAll()
		assert.Nil(t, err)
		assert.Equal(t,actualResponse,payload)
	})

	t.Run("should return error", func(t *testing.T) {
		repoCustomerMock.EXPECT().GetAll().Return(defaultAllCustomerModel,errCustomerInternal)
		response, actualErr := serviceCustomer.GetAll()
		assert.Equal(t,response,defaultAllCustomerModel)
		assert.Equal(t,actualErr,errCustomerInternal)
	})
}
