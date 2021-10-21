package routes

import (
	"AntarJemput-Be-C/handlers"
	"AntarJemput-Be-C/middleware"

	"github.com/gofiber/fiber/v2"
)

type CustomerRoute struct {
	customerHandler handlers.CustomerHandlerInterface
}

func NewCustomerRoutes(customerHandler handlers.CustomerHandlerInterface) *CustomerRoute {
	return &CustomerRoute{customerHandler: customerHandler}
}

func (cr *CustomerRoute) InitialCustomerRoute(app *fiber.App) {
	app.Post(POST_CUSTOMER, middleware.JWTProtected(), cr.customerHandler.Save)
	app.Get(GETALL_CUSTOMER, middleware.JWTProtected(), cr.customerHandler.GetAll)
	app.Get(GETBYID_CUSTOMER, middleware.JWTProtected(), cr.customerHandler.GetById)

}
