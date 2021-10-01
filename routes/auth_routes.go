package routes

import (
	"AntarJemput-Be-C/handlers"

	"github.com/gofiber/fiber/v2"
)

type AuthsRoutes struct {
	authHandler handlers.UserHandlerInterface
}

func NewAuthRoutes(authHandler handlers.UserHandlerInterface) *AuthsRoutes {
	return &AuthsRoutes{authHandler: authHandler}
}

func (ar *AuthsRoutes) InitialAuthRoutes(app *fiber.App) {
	app.Post(POST_REGISTER, ar.authHandler.Register)
	app.Get(GETALL_USERS, ar.authHandler.GetAll)
	// app.Get(GETBYID_USERS, ar.authHandler.GetById)
}
