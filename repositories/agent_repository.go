package repositories

import (
	"AntarJemput-Be-C/models"

	//	"go.mongodb.org/mongo-driver/x/mongo/driver/address"
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
	SearchAgent(city string) (models.Agents, error)
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

func (ar *AgentRepository) SearchAgent(city string) (models.Agents, error) {
	var agent models.Agents
	var transaction models.Transactions
	//query := `SELECT * FROM agent JOIN transaction ON agent.city = city`
	query := `SELECT * FROM agent LEFT JOIN transaction USING(DistrictId)`
	//query := `SELECT * FROM agent WHERE city =?`

	err := ar.DB.Raw(query, city).Scan(&agent).Error
	if err != nil {
		return agent, err
	}

	if agent.DistrictId == transaction.DistrictId {
		return agent, gorm.ErrRecordNotFound
	}

	return agent, nil
}
