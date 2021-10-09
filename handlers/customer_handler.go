package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerService services.CustomerServiceInterface
}

func NewCustomerHandler(customerService services.CustomerServiceInterface) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

type CustomerHandlerInterface interface {
	Save(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

func (ch *CustomerHandler) Save(c *fiber.Ctx) error {
	customer := &models.Customer{}
	err := c.BodyParser(&customer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	
	response, err := ch.customerService.Save(customer)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  200,
		"message": "Success",
		"data":    response,
	})
}

func (ch *CustomerHandler) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response, err := ch.customerService.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    response,
	})
}

func (ch *CustomerHandler) GetAll(c *fiber.Ctx) error {
	customers, err := ch.customerService.GetAll()
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    customers,
	})
}
