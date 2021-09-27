package routes

import (
	"AntarJemput-Be-C/handlers"

	"github.com/gofiber/fiber/v2"
)

type RoleRoutes struct {
	roleHandler handlers.RoleHandlerInterface
}

func NewRoleRoutes(roleHandler handlers.RoleHandlerInterface) *RoleRoutes {
	return &RoleRoutes{roleHandler: roleHandler}
}

func (roleRoutes *RoleRoutes) InitialRoleRoutes(app *fiber.App) {
	app.Post("/role", roleRoutes.roleHandler.Save)
	app.Get("/role", roleRoutes.roleHandler.GetAll)
	app.Get("/role/:id", roleRoutes.roleHandler.GetById)
}
