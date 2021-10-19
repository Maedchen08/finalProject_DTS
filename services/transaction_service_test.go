package services

import (
	mocks "AntarJemput-Be-C/mocks/repository"
	"AntarJemput-Be-C/models"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
)

var (
	mockCtrl                 *gomock.Controller
	repoMock                 *mocks.MockTransactionRepoInterface
	service                  *TransactionService
	errInternal              = errors.New("unexpected system error")
	defaultInt               = 0
	defaultTransactionModels models.Transactions
	defaultRatingAgentsModel models.AgentRating
	defaultAllTransactions   []models.Transactions
)

func InitMockRepo(t *testing.T) {
	mockCtrl = gomock.NewController(t)
	repoMock = mocks.NewMockTransactionRepoInterface(mockCtrl)
	service = NewTransactionService(repoMock)
}

func TestTransactionService_Save(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	expectedId := 1
	trnsac := &models.Transactions{
		TypeTransactionId: 1,
		CustomerId:        1,
		AgentId:           1,
		Address:           "Jalan H. Dahlan",
		ProvinceId:        32,
		RegencyId:         3204,
		DistrictId:        320490,
		Amount:            1000000,
		StatusTransaction: 3,
		Rating:            5,
		RatingComment:     "Good Service",
	}

	expectedResponse := trnsac
	expectedResponse.Id = expectedId

	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().Save(gomock.Any()).Return(expectedId, nil)

		actualResponse, err := service.Save(trnsac)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().Save(gomock.Any()).Return(defaultInt, errInternal)
		respon, actualErr := service.Save(trnsac)
		assert.Nil(t, respon)
		assert.Equal(t, actualErr, errInternal)
	})

}

func TestTransactionService_GetById(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	inputId := 1
	expectedResponse := models.Transactions{
		TypeTransactionId: 1,
		CustomerId:        1,
		AgentId:           1,
		Address:           "Jalan H. Dahlan",
		ProvinceId:        32,
		RegencyId:         3204,
		DistrictId:        320490,
		Amount:            1000000,
		StatusTransaction: 3,
		Rating:            5,
		RatingComment:     "Good Service",
	}

	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(expectedResponse, nil)

		actualResponse, err := service.GetById(inputId)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(defaultTransactionModels, errInternal)

		response, actualErr := service.GetById(inputId)
		assert.Equal(t, response, defaultTransactionModels)
		assert.Equal(t, actualErr, errInternal)
	})
}


func TestTransactionService_GetAll(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	expectedResponse := []models.Transactions{
		{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jalan H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        320490,
			Amount:            1000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Good Service",
		},
		{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jalan H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        320490,
			Amount:            1000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Good Service",
		},
	}

	t.Run("Should return success", func(t *testing.T) {
		repoMock.EXPECT().GetAll().Return(expectedResponse, nil)

		actualResponse, err := service.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().GetAll().Return(defaultAllTransactions, errInternal)

		response, actualErr := service.GetAll()
		assert.Equal(t, response, defaultAllTransactions)
		assert.Equal(t, actualErr, errInternal)
	})
}

func TestTransactionService_DeleteTransaction(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	inputId := 1
	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().DeleteTransaction(gomock.Any()).Return(nil)

		err := service.DeleteTransaction(inputId)
		assert.Nil(t, err)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().DeleteTransaction(gomock.Any()).Return(errInternal)

		actualErr := service.DeleteTransaction(inputId)
		assert.Equal(t, actualErr, errInternal)
	})
}

func TestTransactionService_ChangesConfirmed(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	payload := models.Transactions{
		TypeTransactionId: 1,
		CustomerId:        1,
		AgentId:           1,
		Address:           "Jalan H. Dahlan",
		ProvinceId:        32,
		RegencyId:         3204,
		DistrictId:        320490,
		Amount:            1000000,
		StatusTransaction: 1,
		Rating:            5,
		RatingComment:     "Good Service",
	}

	expectedResponse := payload
	expectedResponse.ID = 1
	t.Run("should return success change status confirmed", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(expectedResponse, nil)
		repoMock.EXPECT().ChangeConfirmed(gomock.Any()).Return(nil)

		actualResponse, err := service.ChangesConfirmed(&payload)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("Should return error change status confirmed", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(defaultTransactionModels, errInternal)
		repoMock.EXPECT().ChangeConfirmed(gomock.Any()).Return(nil)

		response, actualErr := service.ChangesConfirmed(&payload)
		assert.Equal(t, response, defaultTransactionModels)
		assert.Equal(t, actualErr, errInternal)
	})
}

func TestTransactionService_ChangesRejected(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	payload := models.Transactions{
		TypeTransactionId: 1,
		CustomerId:        1,
		AgentId:           1,
		Address:           "Jalan H. Dahlan",
		ProvinceId:        32,
		RegencyId:         3204,
		DistrictId:        320490,
		Amount:            1000000,
		StatusTransaction: 1,
		Rating:            5,
		RatingComment:     "Good Service",
	}

	expectedResponse := payload
	expectedResponse.ID = 1
	t.Run("should return success change status reject", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(expectedResponse, nil)
		repoMock.EXPECT().ChangeReject(gomock.Any()).Return(nil)

		actualResponse, err := service.ChangesRejected(&payload)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("Should return error change status ChangeReject", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(defaultTransactionModels, errInternal)
		repoMock.EXPECT().ChangeReject(gomock.Any()).Return(nil)

		response, actualErr := service.ChangesRejected(&payload)
		assert.Equal(t, response, defaultTransactionModels)
		assert.Equal(t, actualErr, errInternal)
	})
}

func TestTransactionService_ChangesDone(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	payload := models.Transactions{
		TypeTransactionId: 1,
		CustomerId:        1,
		AgentId:           1,
		Address:           "Jalan H. Dahlan",
		ProvinceId:        32,
		RegencyId:         3204,
		DistrictId:        320490,
		Amount:            1000000,
		StatusTransaction: 1,
		Rating:            5,
		RatingComment:     "Good Service",
	}

	expectedResponse := payload
	expectedResponse.ID = 1
	t.Run("should return success change status done", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(expectedResponse, nil)
		repoMock.EXPECT().ChangeDone(gomock.Any()).Return(nil)

		actualResponse, err := service.ChangesDone(&payload)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("Should return error change status done", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(defaultTransactionModels, errInternal)
		repoMock.EXPECT().ChangeDone(gomock.Any()).Return(nil)

		response, actualErr := service.ChangesDone(&payload)
		assert.Equal(t, response, defaultTransactionModels)
		assert.Equal(t, actualErr, errInternal)
	})
}

func TestTransactionService_RatingAgent(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	inputId := 1
	payload := models.AgentRating{
		AgentId:   1,
		Total:     4,
		AvgRating: 4,
	}

	t.Run("should return success rating agent byId", func(t *testing.T) {
		repoMock.EXPECT().RatingAgent(gomock.Any()).Return(payload, nil)

		actualResponse, err := service.RatingAgent(inputId)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, payload)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().RatingAgent(gomock.Any()).Return(defaultRatingAgentsModel, errInternal)

		response, actualErr := service.RatingAgent(inputId)
		assert.Equal(t, response, defaultRatingAgentsModel)
		assert.Equal(t, actualErr, errInternal)
	})

}

func TestTransactionService_GetByAgentId(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	inputId := 1
	expectedResponse := []models.Transactions{
		{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jalan H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        320490,
			Amount:            1000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Good Service",
		},
		{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jalan H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        320490,
			Amount:            1000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Good Service",
		},
	}

	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().GetByAgentId(gomock.Any()).Return(expectedResponse, nil)

		actualResponse, err := service.GetByAgentId(inputId)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("Should return error", func(t *testing.T) {
		repoMock.EXPECT().GetByAgentId(gomock.Any()).Return(defaultAllTransactions, errInternal)

		response, actualErr := service.GetByAgentId(inputId)
		assert.Equal(t, response, defaultAllTransactions)
		assert.Equal(t, actualErr, errInternal)
	})
}

func TestTransactionService_RatingTransaction(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	payload := models.Transactions{
		Rating:        5,
		RatingComment: "Good Service",
	}

	expectedResponse := payload
	expectedResponse.ID = 1

	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(expectedResponse, nil)
		repoMock.EXPECT().RatingTransaction(gomock.Any()).Return(nil)

		actualResponse, err := service.RatingTransaction(&payload)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().GetById(gomock.Any()).Return(defaultTransactionModels, errInternal)
		repoMock.EXPECT().RatingTransaction(gomock.Any()).Return(nil)

		response, actualErr := service.RatingTransaction(&payload)
		assert.Equal(t, response, defaultTransactionModels)
		assert.Equal(t, actualErr, errInternal)
	})

}

func TestTransactionService_GetByCustomerId(t *testing.T) {
	InitMockRepo(t)
	defer mockCtrl.Finish()

	inputId := 1
	payload := []models.Transactions{
		{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jalan H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        320490,
			Amount:            1000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Good Service",
		},
		{
			TypeTransactionId: 1,
			CustomerId:        1,
			AgentId:           1,
			Address:           "Jalan H. Dahlan",
			ProvinceId:        32,
			RegencyId:         3204,
			DistrictId:        320490,
			Amount:            1000000,
			StatusTransaction: 3,
			Rating:            5,
			RatingComment:     "Good Service",
		},
	}

	t.Run("should return success", func(t *testing.T) {
		repoMock.EXPECT().GetByCustomerId(gomock.Any()).Return(payload, nil)

		actualResponse, err := service.GetByCustomerId(inputId)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, payload)
	})

	t.Run("should return error", func(t *testing.T) {
		repoMock.EXPECT().GetByCustomerId(gomock.Any()).Return(defaultAllTransactions, errInternal)

		response, actualErr := service.GetByCustomerId(inputId)
		assert.Equal(t, response, defaultAllTransactions)
		assert.Equal(t, actualErr, errInternal)
	})
}
