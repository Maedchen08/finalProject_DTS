package services

import (
	mocksTS "AntarJemput-Be-C/mocks/repository"
	"AntarJemput-Be-C/models"
	"errors"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
)

var (
	mockTSCtrl        *gomock.Controller
	repoTSMock        *mocksTS.MockTypeServiceInterface
	serviceTS         *TSService
	defaultTSInt      = 0
	defaultTSModel    models.TypeServiceTransaction
	defaultAllTsModel []models.TypeServiceTransaction
	errTSInternal     = errors.New("unexpected system error")
)

func InitMockTSRepo(t *testing.T) {
	mockTSCtrl = gomock.NewController(t)
	repoTSMock = mocksTS.NewMockTypeServiceInterface(mockTSCtrl)
	serviceTS = NewSTService(repoTSMock)
}

func TestTSService_Save(t *testing.T) {
	InitMockTSRepo(t)
	defer mockTSCtrl.Finish()

	expectedId := 1
	payload := &models.TypeServiceTransaction{
		Id:                  1,
		TypeTransactionName: "Report",
	}
	expectedResponse := payload
	expectedResponse.Id = expectedId

	t.Run("should return success", func(t *testing.T) {
		repoTSMock.EXPECT().Save(gomock.Any()).Return(expectedId, nil)
		actualResponse, err := serviceTS.Save(payload)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("should return error", func(t *testing.T) {
		repoTSMock.EXPECT().Save(gomock.Any()).Return(defaultTSInt, errTSInternal)
		response, actualErr := serviceTS.Save(payload)
		assert.Nil(t, response)
		assert.Equal(t, actualErr, errTSInternal)
	})
}

func TestTSService_GetById(t *testing.T) {
	InitMockTSRepo(t)
	defer mockTSCtrl.Finish()

	expectedId := 1
	payload := models.TypeServiceTransaction{
		Id:                  1,
		TypeTransactionName: "Report",
	}
	expectedResponse := payload
	expectedResponse.Id = expectedId

	t.Run("should return success", func(t *testing.T) {
		repoTSMock.EXPECT().GetById(gomock.Any()).Return(payload, nil)
		actualResponse, err := serviceTS.GetById(expectedId)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("should return error", func(t *testing.T) {
		repoTSMock.EXPECT().GetById(gomock.Any()).Return(defaultTSModel, errTSInternal)
		response, actualErr := serviceTS.GetById(expectedId)
		assert.Equal(t, response, defaultTSModel)
		assert.Equal(t, actualErr, errTSInternal)
	})
}

func TestTSService_GetAll(t *testing.T) {
	InitMockTSRepo(t)
	defer mockTSCtrl.Finish()

	payload := []models.TypeServiceTransaction{
		{
			Id:                  1,
			TypeTransactionName: "Report",
		},
		{
			Id:                  2,
			TypeTransactionName: "Report1",
		},
	}


	t.Run("should return success", func(t *testing.T) {
		repoTSMock.EXPECT().GetAll().Return(payload, nil)
		actualResponse, err := serviceTS.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, payload)
	})

	t.Run("should return error", func(t *testing.T) {
		repoTSMock.EXPECT().GetAll().Return(defaultAllTsModel, errTSInternal)
		response, actualErr := serviceTS.GetAll()
		assert.Equal(t, response, defaultAllTsModel)
		assert.Equal(t, actualErr, errTSInternal)
	})
}
