package services

import (
	mocksAgent "AntarJemput-Be-C/mocks/repository"
	"AntarJemput-Be-C/models"
	"errors"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
)

var (
	mockAgentCtrl         *gomock.Controller
	repoAgentMock         *mocksAgent.MockAgentRepoInterface
	serviceAgent          *AgentService
	errAgentInternal      = errors.New("unexpected system error")
	defaultAgentInt       = 0
	defaultAgentModel     models.Agents
	defaultAllAgentsModel []models.Agents
)

func InitMockAgentRepo(t *testing.T) {
	mockAgentCtrl = gomock.NewController(t)
	repoAgentMock = mocksAgent.NewMockAgentRepoInterface(mockAgentCtrl)
	serviceAgent = NewAgentService(repoAgentMock)
}

func TestAgentService_Save(t *testing.T) {
	InitMockAgentRepo(t)
	defer mockAgentCtrl.Finish()

	expectedId := 1
	payload := &models.Agents{
		Id:         1,
		Name:       "Bambang",
		Address:    "Jl H. Dahlan no.75",
		ProvinceId: 32,
		RegencyId:  3204,
		DistrictId: 3204190,
		NoWa:       "089802839411",
	}

	expectedResponse := payload
	expectedResponse.Id = expectedId

	t.Run("should return success", func(t *testing.T) {
		repoAgentMock.EXPECT().Save(gomock.Any()).Return(expectedId, nil)
		actualResponse, err := serviceAgent.Save(payload)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, expectedResponse)
	})

	t.Run("should return error", func(t *testing.T) {
		repoAgentMock.EXPECT().Save(gomock.Any()).Return(defaultAgentInt, errAgentInternal)
		respon, actualErr := serviceAgent.Save(payload)
		assert.Nil(t, respon)
		assert.Equal(t, actualErr, errAgentInternal)
	})
}

func TestAgentService_GetById(t *testing.T) {
	InitMockAgentRepo(t)
	defer mockAgentCtrl.Finish()

	inputId := 1
	payload := models.Agents{
		Id:         1,
		Name:       "Bambang",
		Address:    "Jl H. Dahlan no.75",
		ProvinceId: 32,
		RegencyId:  3204,
		DistrictId: 3204190,
		NoWa:       "089802839411",
	}

	t.Run("should return success", func(t *testing.T) {
		repoAgentMock.EXPECT().GetAgentId(gomock.Any()).Return(payload, nil)
		actualResponse, err := serviceAgent.GetById(inputId)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, payload)
	})

	t.Run("should return error", func(t *testing.T) {
		repoAgentMock.EXPECT().GetAgentId(gomock.Any()).Return(defaultAgentModel, errAgentInternal)
		response, actualErr := serviceAgent.GetById(inputId)
		assert.Equal(t, response, defaultAgentModel)
		assert.Equal(t, actualErr, errAgentInternal)
	})
}

func TestAgentService_GetAll(t *testing.T) {
	InitMockAgentRepo(t)
	defer mockAgentCtrl.Finish()

	payload := []models.Agents{
		{
			Id:         1,
			Name:       "Bambang",
			Address:    "Jl H. Dahlan no.75",
			ProvinceId: 32,
			RegencyId:  3204,
			DistrictId: 3204190,
			NoWa:       "089802839411",
		},
		{
			Id:         2,
			Name:       "Bambang",
			Address:    "Jl H. Dahlan no.75",
			ProvinceId: 32,
			RegencyId:  3204,
			DistrictId: 3204190,
			NoWa:       "089802839411",
		},
		{
			Id:         3,
			Name:       "Bambang",
			Address:    "Jl H. Dahlan no.75",
			ProvinceId: 32,
			RegencyId:  3204,
			DistrictId: 3204190,
			NoWa:       "089802839411",
		},
	}

	t.Run("should return success", func(t *testing.T) {
		repoAgentMock.EXPECT().GetAgent().Return(payload, nil)
		actualResponse, err := serviceAgent.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, actualResponse, payload)
	})

	t.Run("should return error", func(t *testing.T) {
		repoAgentMock.EXPECT().GetAgent().Return(defaultAllAgentsModel, errAgentInternal)
		response, actualErr := serviceAgent.GetAll()
		assert.Equal(t, response, defaultAllAgentsModel)
		assert.Equal(t, actualErr, errAgentInternal)
	})
}

func TestAgentService_SearchAgent(t *testing.T) {
	InitMockAgentRepo(t)
	defer mockAgentCtrl.Finish()

	inputId := 1
	payload := []models.Agents{
		{
			Id:         1,
			Name:       "Bambang",
			Address:    "Jl H. Dahlan no.75",
			ProvinceId: 32,
			RegencyId:  3204,
			DistrictId: 3204190,
			NoWa:       "089802839411",
		},
		{
			Id:         2,
			Name:       "Bambang",
			Address:    "Jl H. Dahlan no.75",
			ProvinceId: 32,
			RegencyId:  3204,
			DistrictId: 3204190,
			NoWa:       "089802839411",
		},
		{
			Id:         3,
			Name:       "Bambang",
			Address:    "Jl H. Dahlan no.75",
			ProvinceId: 32,
			RegencyId:  3204,
			DistrictId: 3204190,
			NoWa:       "089802839411",
		},
	}

	t.Run("should return success", func(t *testing.T) {
		repoAgentMock.EXPECT().SearchAgent(gomock.Any()).Return(payload, nil)
		actualResponse,err := repoAgentMock.SearchAgent(inputId)
		assert.Nil(t, err)
		assert.Equal(t, actualResponse,payload)
	})

	t.Run("should return error",func(t *testing.T) {
		repoAgentMock.EXPECT().SearchAgent(gomock.Any()).Return(defaultAllAgentsModel,errAgentInternal)
		response, actualErr := serviceAgent.SearchAgent(inputId)
		assert.Equal(t,response,defaultAllAgentsModel)
		assert.Equal(t, actualErr, errAgentInternal)
	})
}
