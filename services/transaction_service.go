package services

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/repositories"
)

type TransactionService struct {
	transaction repositories.TransactionRepoInterface
}

func NewTransactionService(transaction repositories.TransactionRepoInterface) *TransactionService {
	return &TransactionService{transaction: transaction}
}

type TransactionServiceInterface interface {
	Save(t *models.Transactions) (*models.Transactions, error)
}

func (ts *TransactionService) Save(newTransaction *models.Transactions) (*models.Transactions, error) {
	id, err := ts.transaction.Save(newTransaction)
	if err != nil {
		return nil, err
	}
	newTransaction.Id = id
	return newTransaction, nil
}
