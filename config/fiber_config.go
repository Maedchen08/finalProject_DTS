package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	readTimeoutSecondsCount := GetInt("SERVER_READ_TIMEOUT", 60)
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
