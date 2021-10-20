package middleware

import (
	"AntarJemput-Be-C/config"
	"fmt"

	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   []byte(config.GetString("JWT_SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}
	fmt.Println("Config :", config)
	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
		"mes":   "Status Unauthorized",
	})
}
