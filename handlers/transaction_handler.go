package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"
	"AntarJemput-Be-C/utils"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	transactionService services.TransactionServiceInterface
}

func NewTransactionHandler(typeTransaction services.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{transactionService: typeTransaction}
}

type TransactionHandlerInterface interface {
	Save(c *fiber.Ctx) error
	ChangeConfirmed(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

//Add Transaction
func (th *TransactionHandler) Save(c *fiber.Ctx) error {
	trans := &models.Transactions{}
	err := c.BodyParser(&trans)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
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

//edit status confirmed
func (th *TransactionHandler) ChangeConfirmed(c *fiber.Ctx) error {
	transaction := &models.Transactions{}
	var mysqlErr *mysql.MySQLError
	err1 := c.BodyParser(transaction)
	if err1 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err1.Error(),
		})
	}
	validate := utils.NewValidator()
	err := validate.Struct(transaction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": utils.ValidatorErrors(err),
		})
	}
	response, err := th.transactionService.ChangesConfirmed(transaction)
	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "database not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "data has been update",
		"result":  response,
	})

}

func (th *TransactionHandler) GetAll(c *fiber.Ctx) error {
	transaction, err := th.transactionService.GetAll()
	if err != nil {
		c.Status(403).JSON(err)
	}
	return c.JSON(transaction)
}
