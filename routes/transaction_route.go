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
	app.Post(POST_CONFIRM_TRANSACTION, tr.transactionHandler.ChangeConfirmed)
	app.Get(GETALL_TRANSACTION, tr.transactionHandler.GetAll)
	app.Get(GETBYID_TRANSACTION, tr.transactionHandler.GetById)
	app.Post(POST_REJECT_TRANSACTION, tr.transactionHandler.ChangeRejected)
	app.Post(POST_DONE_TRANSACTION, tr.transactionHandler.ChangeDone)
	app.Delete(DELETE_TRANSACTION,tr.transactionHandler.DeleteTransaction)
	app.Get(GETBYID_TRANSACTION_AGENT, tr.transactionHandler.GetByAgentId)
	app.Get(GETBYID_TRANSACTION_CUSTOMER, tr.transactionHandler.GetByCustomerId)
	app.Post(POST_RATING_TRANSACTION, tr.transactionHandler.RatingTransaction)
	app.Get(GETRATING_AGENT, tr.transactionHandler.RatingAgent)

}
