package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SecretKey = "secret"

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *AuthHandler) Register(c *fiber.Ctx) error {

	var data models.Users

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := HashPassword(data.Password)
	user := models.Users{
		RoleId:     data.RoleId,
		CustomerId: data.CustomerId,
		AgentId:    data.AgentId,
		Username:   data.Username,
		Password:   password,
	}

	a.authService.Register(&user)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		// "status":  201,
		"message": "Success",
		"data":    user,
	})
}

func (a *AuthHandler) Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Get POST
	username := data["username"]
	password := data["password"]

	respone, _ := a.authService.Login(username)

	// handle error
	if uint(respone.Id) == 0 { //default Id when return nil
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found!",
		})
	}

	// match password
	if check := CheckPasswordHash(password, respone.Password); !check {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password!",
		})
	}

	claims := jwt.MapClaims{}
	claims["username"] = respone.Username
	claims["login_as"] = respone.RoleId.EnumIndex()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() //1 day

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	fmt.Println("Print respone ", respone)
	fmt.Println("Print role id enum", respone.RoleId.EnumIndex())
	fmt.Println("Print role id", respone.RoleId)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "error when logging in !",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status":   200,
		"message":  "Login Succeeded",
		"token":    t,
		"username": respone.Username,
		"Id":       respone.CustomerId,
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
