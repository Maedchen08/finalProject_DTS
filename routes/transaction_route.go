package routes

import (
	"AntarJemput-Be-C/handlers"

	"github.com/gofiber/fiber/v2"
)

type TransactionRoutes struct {
	transactionHandler handlers.TransactionHandlerInterface
}

func NewTransRoutes(transHandler handlers.TransactionHandlerInterface) *TransactionRoutes {
	return &TransactionRoutes{transactionHandler: transHandler}
}

func (tr *TransactionRoutes) InitialTransactionRoutes(app *fiber.App) {
	app.Post(POST_TRANSACTION, tr.transactionHandler.Save)
}
