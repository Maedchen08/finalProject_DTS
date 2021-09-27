package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	roleService services.RoleServiceInterface
}

func NewRoleHandler(roleService services.RoleServiceInterface) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

type RoleHandlerInterface interface {
	Save(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

func (rh *RoleHandler) Save(c *fiber.Ctx) error {
	role := &models.Role{}
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  true,
			"message": err.Error(),
		})
	}

	response, err := rh.roleService.Save(role)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  false,
		"message": "Success",
		"data":    response,
	})
}

func (rh *RoleHandler) GetAll(c *fiber.Ctx) error {
	roles, err := rh.roleService.GetAll()
	if err != nil {
		c.Status(403).JSON(err)
	}
	return c.JSON(roles)
}

func (rh *RoleHandler) GetById(c *fiber.Ctx) error {
	id,_ := strconv.Atoi(c.Params("id"))
	response, err := rh.roleService.GetById(id)
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
