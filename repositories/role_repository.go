package repositories

import (
	"AntarJemput-Be-C/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		DB: db,
	}
}

//Role: represent the customers repository contract
type RoleRepoInterface interface{
	Save(models.Role) (models.Role, error)
	GetRole() ([]models.Role,error)
	GetRoleId(id int) (models.Role,error)
}

func (roleRepository *RoleRepository) Save(role models.Role) (models.Role, error) {
	trx := roleRepository.DB.Create(&role)
	return role, trx.Error
}

func (roleRepository *RoleRepository) GetRole() ([]models.Role, error) {
	var role []models.Role
	findRole := roleRepository.DB.Find(&role)
	return role, findRole.Error
}

func (roleRepository *RoleRepository) GetRoleId(id int) (models.Role, error) {
	var role models.Role
	query := `SELECT id,name FROM role WHERE id =?`
	err := roleRepository.DB.Raw(query,id).Scan(&role).Error
	
	if err != nil {
		return role, gorm.ErrRecordNotFound
	}
	return role, nil
}



