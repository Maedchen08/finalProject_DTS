package routes

import (
	"AntarJemput-Be-C/handlers"

	"github.com/gofiber/fiber/v2"
)

type CustomerRoute struct {
	customerHandler handlers.CustomerHandlerInterface
}

func NewCustomerRoutes(customerHandler handlers.CustomerHandlerInterface) *CustomerRoute {
	return &CustomerRoute{customerHandler: customerHandler}
}

func (cr *CustomerRoute) InitialCustomerRoute(app *fiber.App) {
	app.Post(POST_CUSTOMER, cr.customerHandler.Save)
	app.Get(GETALL_CUSTOMER, cr.customerHandler.GetAll)
	app.Get(GETBYID_CUSTOMER, cr.customerHandler.GetById)
}
