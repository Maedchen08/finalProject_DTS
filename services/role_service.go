package services

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/repositories"
)

type RoleService struct {
	roleRepository repositories.RoleRepoInterface
}

func NewRoleService(roleRepository repositories.RoleRepoInterface) *RoleService {
	return &RoleService{roleRepository: roleRepository}
}

type RoleServiceInterface interface {
	Save(role *models.Role) (*models.Role, error)
	GetAll() ([]models.Role, error)
	GetById(id int) (models.Role, error)
}

func(roleService *RoleService) Save(role *models.Role) (*models.Role, error) {
	id, err := roleService.roleRepository.Save(role)
	if err != nil {
		return nil, err
	}
	role.Id = id
	return role, nil
}

func (roleService *RoleService) GetAll() ([]models.Role, error){
	roles,err := roleService.roleRepository.GetRole()
	return roles,err
}

func (roleService *RoleService) GetById(id int) (models.Role, error){
	roles,err := roleService.roleRepository.GetRoleId(id)
	if err != nil {
		return roles,err
	}
	return roles,nil
}