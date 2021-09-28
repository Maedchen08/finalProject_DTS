package repositories

import (
	"AntarJemput-Be-C/models"

	"gorm.io/gorm"
)

type AgentRepository struct {
	DB *gorm.DB
}

func NewAgentRepository(db *gorm.DB) *AgentRepository {
	return &AgentRepository{
		DB: db,
	}
}

//represent the agent repository contract
type AgentRepoInterface interface {
	Save(agent *models.Agents) (int, error)
	GetAgent() ([]models.Agents, error)
	GetAgentId(id int) (models.Agents, error)
}

func (ar *AgentRepository) Save(agent *models.Agents) (int, error) {
	err := ar.DB.Create(agent).Error
	if err != nil {
		return agent.Id, err
	}
	return agent.Id, nil
}

func (ar *AgentRepository) GetAgent() ([]models.Agents, error) {
	var agents []models.Agents
	findAgent := ar.DB.Find(&agents)
	return agents, findAgent.Error
}

func (ar *AgentRepository) GetAgentId(id int) (models.Agents, error) {
	var agent models.Agents
	query := `SELECT * FROM agent WHERE id =?`
	err := ar.DB.Raw(query, id).Scan(&agent).Error

	if err != nil {
		return agent, err
	}
	if agent.Id == 0 {
		return agent, gorm.ErrRecordNotFound
	}
	return agent, nil
}