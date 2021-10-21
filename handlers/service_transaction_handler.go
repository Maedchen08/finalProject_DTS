package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type STransacHandler struct {
	sTransactionService services.STSInterface
}

func NewSTHandler(sts services.STSInterface) *STransacHandler {
	return &STransacHandler{sTransactionService: sts}
}

type STHInterface interface {
	Save(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

func (sth *STransacHandler) Save(c *fiber.Ctx) error {
	st := &models.ServiceTransaction{}
	err := c.BodyParser(&st)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	response, err := sth.sTransactionService.Save(st)
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

func (sth *STransacHandler) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response, err := sth.sTransactionService.GetById(id)
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

func (sth *STransacHandler) GetAll(c *fiber.Ctx) error {
	st, err := sth.sTransactionService.GetAll()
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    st,
	})
}
