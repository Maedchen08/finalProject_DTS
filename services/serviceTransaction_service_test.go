package services

import (
	mocksST "AntarJemput-Be-C/mocks/repository"
	"AntarJemput-Be-C/models"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
)

var (
	mocksSTCtrl       *gomock.Controller
	repoSTMock        *mocksST.MockServiceTransacInterface
	serviceST         *STransactionService
	defaultSTInt      = 0
	defaultSTModel    models.ServiceTransaction
	defaultAllSTModel []models.ServiceTransaction
	errSTInternal     = errors.New("unexpected system error")
)

func InitMockSTRepo(t *testing.T) {
	mocksSTCtrl = gomock.NewController(t)
	repoSTMock = mocksST.NewMockServiceTransacInterface(mocksSTCtrl)
	serviceST = NewSTransactionService(repoSTMock)
}

func TestSTransactionService_Save(t *testing.T) {
	InitMockSTRepo(t)
	defer mocksSTCtrl.Finish()

	expectedId := 1
	payload := &models.ServiceTransaction{
		Id:          1,
		ServiceName: "Setoran Tunai",
	}

	expectedResponse := payload
	expectedResponse.Id = expectedId

	t.Run("should return success", func(t *testing.T) {
		repoSTMock.EXPECT().Save(gomock.Any()).Return(expectedId, nil)
		actualResponse, err := serviceST.Save(payload)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("should return error", func(t *testing.T) {
		repoSTMock.EXPECT().Save(gomock.Any()).Return(defaultSTInt, errSTInternal)
		response, actualErr := serviceST.Save(payload)
		assert.Nil(t, response)
		assert.Equal(t, actualErr, errSTInternal)
	})

}

func TestSTransactionService_GetAll(t *testing.T) {
	InitMockSTRepo(t)
	defer mocksSTCtrl.Finish()

	payload := []models.ServiceTransaction{
		{
			Id:          1,
			ServiceName: "Setoran Tunai",
		},
		{
			Id:          2,
			ServiceName: "Pinjaman Online",
		},
	}

	t.Run("should return success", func(t *testing.T) {
		repoSTMock.EXPECT().GetAll().Return(payload, nil)
		actualResponse, err := serviceST.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, payload)
	})

	t.Run("should return error", func(t *testing.T) {
		repoSTMock.EXPECT().GetAll().Return(defaultAllSTModel, errSTInternal)
		response, actualErr := serviceST.GetAll()
		assert.Equal(t, response, defaultAllSTModel)
		assert.Equal(t, actualErr, errSTInternal)

	})
}

func TestSTransactionService_GetById(t *testing.T) {
	InitMockSTRepo(t)
	defer mocksSTCtrl.Finish()

	expectedId := 1
	payload := models.ServiceTransaction{
		Id:          1,
		ServiceName: "Setoran Tunai",
	}

	t.Run("should return success", func(t *testing.T) {
		repoSTMock.EXPECT().GetById(gomock.Any()).Return(payload, nil)
		actualResponse, err := serviceST.GetById(expectedId)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, payload)
	})

	t.Run("should return error", func(t *testing.T) {
		repoSTMock.EXPECT().GetById(gomock.Any()).Return(defaultSTModel, errSTInternal)
		response, actualErr := serviceST.GetById(expectedId)
		assert.Equal(t, response,defaultSTModel)
		assert.Equal(t, actualErr, errSTInternal)
	})
}
