package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"
	"fmt"
	"strconv"
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

func (a *AuthHandler) Register(c *fiber.Ctx) error {

	var data models.Users

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
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

	var user models.Users
	username := data["username"]
	password := data["password"]
	// a.DB.Where("username = ?", data["username"]).First(&user)
	respone, _ := a.authService.Login(username, password)
	fmt.Println(user)

	// handle error
	if uint(respone.Id) == 0 { //default Id when return nil
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found!",
		})
	}
	// repassword := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	// fmt.Print(repassword)

	// match password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password!",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(respone.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "error when logging in !",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Login Succeeded",
		"token":   token,
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
