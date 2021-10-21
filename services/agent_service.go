package services

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/repositories"

)



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



