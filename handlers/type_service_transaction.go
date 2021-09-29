package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TypeServiceHandler struct {
	typeService services.TSSInterface
}

func NewTSHandler(typeService services.TSSInterface) *TypeServiceHandler {
	return &TypeServiceHandler{typeService: typeService}
}

type TSHandlerInterface interface {
	Save(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

func (tsh *TypeServiceHandler) Save(c *fiber.Ctx) error {
	ts := &models.TypeServiceTransaction{}
	err := c.BodyParser(&ts)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	response, err := tsh.typeService.Save(ts)

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

func (tsh *TypeServiceHandler) GetAll(c *fiber.Ctx) error {
	ts, err := tsh.typeService.GetAll()
	if err != nil {
		c.Status(403).JSON(err)
	}
	return c.JSON(ts)
}

func (tsh *TypeServiceHandler) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response, err := tsh.typeService.GetById(id)
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
