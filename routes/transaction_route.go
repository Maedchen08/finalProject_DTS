package routes

import (
	"AntarJemput-Be-C/handlers"
	"AntarJemput-Be-C/middleware"

	"github.com/gofiber/fiber/v2"
)

type TransactionRoutes struct {
	transactionHandler handlers.TransactionHandlerInterface
}

func NewTransRoutes(transHandler handlers.TransactionHandlerInterface) *TransactionRoutes {
	return &TransactionRoutes{transactionHandler: transHandler}
}

func (tr *TransactionRoutes) InitialTransactionRoutes(app *fiber.App) {
	app.Post(POST_TRANSACTION, middleware.JWTProtected(), tr.transactionHandler.Save)
	app.Post(POST_CONFIRM_TRANSACTION, middleware.JWTProtected(), tr.transactionHandler.ChangeConfirmed)
	app.Get(GETALL_TRANSACTION, middleware.JWTProtected(), tr.transactionHandler.GetAll)
	app.Get(GETBYID_TRANSACTION, middleware.JWTProtected(), tr.transactionHandler.GetById)
	app.Post(POST_REJECT_TRANSACTION, middleware.JWTProtected(), tr.transactionHandler.ChangeRejected)
	app.Post(POST_DONE_TRANSACTION, middleware.JWTProtected(), tr.transactionHandler.ChangeDone)
	app.Delete(DELETE_TRANSACTION, middleware.JWTProtected(), tr.transactionHandler.DeleteTransaction)
	app.Get(GETBYID_TRANSACTION_AGENT, middleware.JWTProtected(), tr.transactionHandler.GetByAgentId)
	app.Get(GETBYID_TRANSACTION_CUSTOMER, middleware.JWTProtected(), tr.transactionHandler.GetByCustomerId)
	app.Post(POST_RATING_TRANSACTION, middleware.JWTProtected(), tr.transactionHandler.RatingTransaction)
	app.Get(GETRATING_AGENT, middleware.JWTProtected(), tr.transactionHandler.RatingAgent)

}
