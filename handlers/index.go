package handlers

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    fiber.Map{
			"nama_kelompok":"backend_kelompok_c",
			"projek":"AntarJemput-Be-C",
		},
	})
}