package repositories

import (
	"AntarJemput-Be-C/models"
	//"fmt"

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
	SearchAgent(s int) ([]models.Agents, error)
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

func (sa *AgentRepository) SearchAgent(s int) ([]models.Agents, error) {
	var agents []models.Agents
	findAgent := sa.DB.Where("agent_district_id = ?", s).Find(&agents)
	return agents, findAgent.Error

	// var agents []models.Agents
	// //findAgent := sa.DB.Joins("JOIN Transaction on transaction_district_id = agent_district_id").Where("transaction.transaction_district_id = ?", s).Find(&agents)
	// findAgent := sa.DB.Table("Transaction").Where("transaction.transaction_district_id = agent_district_id = ?", s).Select("agent_name").Find(&agents)
	// return agents, findAgent.Error

	// var agent models.Agents
	// var transaction models.Transactions
	// query := `SELECT id, agent_name, address, no_wa from agent JOIN transaction WHERE agent_district_id = s`
	// err := sa.DB.Raw(query, s).Scan(&agent).Error

	// if err != nil {
	// 	return agent, err
	// }
	// if transaction.DistrictId == 0 {
	// 	return agent, gorm.ErrRecordNotFound
	// }
	// return agent, nil

}


