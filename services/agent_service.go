package services

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/repositories"
	"math"
)

const km = float64(1)
const earthRadius = float64(6371)

type AgentService struct {
	agentRepository repositories.AgentRepoInterface
}

func NewAgentService(agentRepository repositories.AgentRepoInterface) *AgentService {
	return &AgentService{agentRepository: agentRepository}
}

type AgentServiceInterface interface {
	Save(agent *models.Agents) (*models.Agents, error)
	GetAll() ([]models.Agents, error)
	GetById(id int) (models.Agents, error)
	SearchAgent(s int) ([]models.Agents, error)
}

func (as *AgentService) Save(agent *models.Agents) (*models.Agents, error) {
	id, err := as.agentRepository.Save(agent)
	if err != nil {
		return nil, err
	}
	agent.Id = id
	return agent, nil
}

func (as *AgentService) GetById(id int) (models.Agents, error) {
	agent, err := as.agentRepository.GetAgentId(id)
	if err != nil {
		return agent, err
	}
	return agent, nil
}

func (as *AgentService) GetAll() ([]models.Agents, error) {
	agents, err := as.agentRepository.GetAgent()
	if err != nil {
		return agents, err
	}
	return agents, nil
}

func (as *AgentService) SearchAgent(s int) ([]models.Agents, error) {
	agent, err := as.agentRepository.SearchAgent(s)
	if err != nil {
		return agent, err
	}

	return agent, nil
}


// func (as *AgentService) SearchWithPicker(s int) ([]models.Agents, error){
	

// }

func Haversine(lonFrom float64, latFrom float64, lonTo float64, latTo float64) (distance float64) {
	var deltaLat = (latTo - latFrom) * (math.Pi / 180)
	var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)
	
	var a = math.Sin(deltaLat / 2) * math.Sin(deltaLat / 2) + 
		math.Cos(latFrom * (math.Pi / 180)) * math.Cos(latTo * (math.Pi / 180)) *
		math.Sin(deltaLon / 2) * math.Sin(deltaLon / 2)
	var c = 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))
	
	distance = earthRadius * c
	
	return
}

