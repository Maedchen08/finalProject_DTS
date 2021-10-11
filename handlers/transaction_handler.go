package handlers

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/services"
	"AntarJemput-Be-C/utils"
	"errors"
	"strconv"

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
	ChangeRejected(c *fiber.Ctx) error
	ChangeDone(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	DeleteTransaction(c *fiber.Ctx) error
	//GetByIdAgen(c *fiber.Ctx) error

	GetByAgentId(c *fiber.Ctx) error
	GetByCustomerId(c *fiber.Ctx) error
	RatingTransaction(c *fiber.Ctx) error
	RatingAgent(c *fiber.Ctx) error
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

	if trans.Amount < 50000 && trans.Amount <= 10000000 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "please fill amount more than 50000 and less than 10000000",
		})
	}
	if trans.AgentId == 0 {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"status":  202,
			"message": "Data agen tidak ditemukan",
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

//edit status REJECTED
func (th *TransactionHandler) ChangeRejected(c *fiber.Ctx) error {
	transaction := &models.Transactions{}
	err1 := c.BodyParser(transaction)
	if err1 != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": err1.Error(),
		})
	}

	response, err := th.transactionService.ChangesRejected(transaction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   400,
			"message": "data not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "data has been update",
		"result":  response,
	})

}

//edit status confirmed
func (th *TransactionHandler) ChangeConfirmed(c *fiber.Ctx) error {
	transaction := &models.Transactions{}
	err1 := c.BodyParser(transaction)
	if err1 != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": err1.Error(),
		})
	}

	response, err := th.transactionService.ChangesConfirmed(transaction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   400,
			"message": "data not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "data has been update",
		"result":  response,
	})

}

//edit status done
func (th *TransactionHandler) ChangeDone(c *fiber.Ctx) error {
	transaction := &models.Transactions{}
	err1 := c.BodyParser(transaction)
	if err1 != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": err1.Error(),
		})
	}

	response, err := th.transactionService.ChangesDone(transaction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   400,
			"message": "data not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
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

func (th *TransactionHandler) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response, err := th.transactionService.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    response,
	})
}

//delete
func (th *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := th.transactionService.DeleteTransaction(id)
	if err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		// 		"error":true,
		// 		"message":"data not found",

		// 	})
		// }

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success delete transaction",
		"data": fiber.Map{
			"id_transaction": id,
		},
	})
}

func (th *TransactionHandler) GetByAgentId(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response, err := th.transactionService.GetByAgentId(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    response,
	})
}

func (th *TransactionHandler) GetByCustomerId(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response, err := th.transactionService.GetByCustomerId(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    response,
	})
}

func (th *TransactionHandler) RatingTransaction(c *fiber.Ctx) error {
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
	response, err := th.transactionService.RatingTransaction(transaction)
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

func (th *TransactionHandler) RatingAgent(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response, err := th.transactionService.RatingAgent(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  404,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    response,
	})
}