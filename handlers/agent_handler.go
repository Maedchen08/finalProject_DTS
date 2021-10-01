package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AgentHandler struct {
	agentService services.AgentServiceInterface
}

func NewAgentHandler(agentService services.AgentServiceInterface) *AgentHandler {
	return &AgentHandler{agentService: agentService}
}

type AgentHandlerInterface interface {
	Save(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

func (ah *AgentHandler) Save(c *fiber.Ctx) error {
	agent := &models.Agents{}
	err := c.BodyParser(&agent)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	response, err := ah.agentService.Save(agent)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"message": "Success",
		"data":    response,
	})
}

func (ah *AgentHandler) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response, err := ah.agentService.GetById(id)
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

func (ah *AgentHandler) GetAll(c *fiber.Ctx) error {
	agents, err := ah.agentService.GetAll()
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    agents,
	})
}
