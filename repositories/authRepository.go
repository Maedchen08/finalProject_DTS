package repositories

import (
	"AntarJemput-Be-C/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

//represent the users repository contract
type AuthRepoInterface interface {
	Register(users *models.Users) (int, error)
	GetUser() ([]models.Users, error)
	GetUserId(id int) (models.Users, error)
	Login(username, password string) (models.Users, error)
}

func (ar *AuthRepository) Register(users *models.Users) (int, error) {
	err := ar.DB.Create(users).Error
	if err != nil {
		return users.Id, err
	}
	return users.Id, nil
}

func (ar *AuthRepository) GetUser() ([]models.Users, error) {
	var users []models.Users
	findUser := ar.DB.Find(&users)
	return users, findUser.Error
}

func (ar *AuthRepository) Login(username, password string) (models.Users, error) {
	var users models.Users
	query := `SELECT * FROM Users WHERE username = ?`

	err := ar.DB.Raw(query, username).Scan(&users).Error

	if err != nil {
		return users, err
	}
	if err := bcrypt.CompareHashAndPassword(users.Password, []byte(password)); err != nil {
		// ar.Status(fiber.StatusBadRequest)
		return users, nil
	}

	// handle error
	// if user.Id == 0 { //default Id when return nil
	// 	c.Status(fiber.StatusNotFound)
	// 	return c.JSON(fiber.Map{
	// 		"message": "User not found!",
	// 	})
	// }

	return users, nil
}

func (ar *AuthRepository) GetUserId(id int) (models.Users, error) {
	var users models.Users
	query := `SELECT * FROM Users WHERE id =?`
	err := ar.DB.Raw(query, id).Scan(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
