package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// const SecretKey = "secret"

type AuthHandler struct {
	authService services.AuthServiceInterface
	DB          *gorm.DB
}

func NewAuthHandler(authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type UserHandlerInterface interface {
	Register(c *fiber.Ctx) error
	// GetById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

func (a *AuthHandler) Register(c *fiber.Ctx) error {
	// var data map[string]interface{}
	// var input models.Users
	users := &models.Users{}
	err := c.BodyParser(&users)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	// Hashing the password with the default cost of 10
	// plaintext := []byte("textstring")
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)

	// respone := models.Users{
	// 	CustomerId: 1,
	// }

	a.authService.Register(users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		// "status":  201,
		"message": "Success",
		// "data":    users,
	})
}

func (a *AuthHandler) Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// users := &models.Users{}

	// findLogin := a.DB.Where("username = ?").First(&users)
	// findLogin, _ := a.authService.Login(data["username"])

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    data,
	})
}

func (ah *AuthHandler) GetAll(c *fiber.Ctx) error {
	users, err := ah.authService.GetAll()
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    users,
	})
}
