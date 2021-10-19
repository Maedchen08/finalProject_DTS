package routes

import (
	"AntarJemput-Be-C/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitialIndexRoutes(app *fiber.App) {
	app.Get("/",handlers.Index)
	
}