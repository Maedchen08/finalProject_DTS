package routes

import (
	"AntarJemput-Be-C/handlers"
	"AntarJemput-Be-C/middleware"

	"github.com/gofiber/fiber/v2"
)

type STransacRoute struct {
	sTHandler handlers.STHInterface
}

func NewSTRoutes(sTHandler handlers.STHInterface) *STransacRoute {
	return &STransacRoute{sTHandler: sTHandler}
}

func (str *STransacRoute) InitialSTRoute(app *fiber.App) {
	app.Post(POST_SERVICE_TRANSACTION, middleware.JWTProtected(), str.sTHandler.Save)
	app.Get(GETALL_SERVICE_TRANSACTION, middleware.JWTProtected(), str.sTHandler.GetAll)
	app.Get(GETBYID_SERVICE_TRANSACTION, middleware.JWTProtected(), str.sTHandler.GetById)
}
