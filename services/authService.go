package services

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/repositories"
)

type AuthService struct {
	authRepository repositories.AuthRepoInterface
}

func NewAuthService(authRepository repositories.AuthRepoInterface) *AuthService {
	return &AuthService{authRepository: authRepository}
}

type AuthServiceInterface interface {
	Register(users *models.Users) (*models.Users, error)
	GetAll() ([]models.Users, error)
	GetById(id int) (models.Users, error)
	// Login(username string) (models.Users, error)
}

func (as *AuthService) Register(users *models.Users) (*models.Users, error) {
	id, err := as.authRepository.Register(users)
	if err != nil {
		return nil, err
	}
	users.Id = id
	return users, nil
}

func (as *AuthService) GetById(id int) (models.Users, error) {
	users, err := as.authRepository.GetUserId(id)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (as *AuthService) GetAll() ([]models.Users, error) {
	users, err := as.authRepository.GetUser()
	if err != nil {
		return users, err
	}
	return users, nil
}

// func (as *AuthService) Login() ([]models.Users, error) {
// 	users, err := as.authRepository.GetUser()
// 	if err != nil {
// 		return users, err
// 	}
// 	return users, nil
// }
