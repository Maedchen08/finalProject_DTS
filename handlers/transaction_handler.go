package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionService services.TransactionServiceInterface
}

func NewTransactionHandler(typeTransaction services.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{transactionService: typeTransaction}
}

type TransactionHandlerInterface interface {
	Save(c *fiber.Ctx) error
}

//Add Transaction
func (th *TransactionHandler) Save(c *fiber.Ctx) error {
	trans := &models.Transactions{}
	err := c.BodyParser(&trans)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400 ,
			"message": err.Error(),
		})
	}
	response, err := th.transactionService.Save(trans)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"message": "Success",
		"data":    response,
	})

}
