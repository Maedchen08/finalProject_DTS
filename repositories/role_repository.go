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

//Role: represent the role repository contract
type RoleRepoInterface interface {
	Save(role *models.Role) (int, error)
	GetRole() ([]models.Role, error)
	GetRoleId(id int) (models.Role, error)
}

func (rp *RoleRepository) Save(role *models.Role) (int, error) {
	err := rp.DB.Create(role).Error
	if err != nil {
		return role.Id, err
	}
	return role.Id, nil
}

func (rp *RoleRepository) GetRole() ([]models.Role, error) {
	var role []models.Role
	findRole := rp.DB.Find(&role)
	return role, findRole.Error
}

func (rp *RoleRepository) GetRoleId(id int) (models.Role, error) {
	var role models.Role
	query := `SELECT id,role_name FROM role WHERE id =?`
	err := rp.DB.Raw(query, id).Scan(&role).Error

	if err != nil {
		return role, err
	}
	if role.Id == 0 {
		return role, gorm.ErrRecordNotFound
	}
	return role, nil
}


