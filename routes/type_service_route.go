package routes

import (
	"AntarJemput-Be-C/handlers"
	"AntarJemput-Be-C/middleware"

	"github.com/gofiber/fiber/v2"
)

type TSRoute struct {
	tSHandler handlers.TSHandlerInterface
}

func NewTSRoutes(tSHandler handlers.TSHandlerInterface) *TSRoute {
	return &TSRoute{tSHandler: tSHandler}
}

func (tSR *TSRoute) InitialTSRoute(app *fiber.App){
	app.Post(POST_TYPE_SERVICE_TRANSACTION, middleware.JWTProtected(),tSR.tSHandler.Save)
	app.Get(GETALL_TYPE_SERVICE_TRANSACTION,middleware.JWTProtected(), tSR.tSHandler.GetAll)
	app.Get(GETBYID_TYPE_SERVICE_TRANSACTION,middleware.JWTProtected(), tSR.tSHandler.GetById)
}