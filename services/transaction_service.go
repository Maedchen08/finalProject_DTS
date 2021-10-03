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
	ChangesConfirmed(t *models.Transactions) (models.Transactions, error)
	GetAll() ([]models.Transactions, error)
}

func (ts *TransactionService) Save(newTransaction *models.Transactions) (*models.Transactions, error) {
	id, err := ts.transaction.Save(newTransaction)
	if err != nil {
		return nil, err
	}
	newTransaction.Id = id
	return newTransaction, nil
}

func (ts *TransactionService) ChangesConfirmed(t *models.Transactions) (trans models.Transactions, err error) {
	err = ts.transaction.ChangeConfirmed(t)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

func (ts *TransactionService) GetAll() ([]models.Transactions, error){
	transac, err := ts.transaction.GetAll()
	return transac, err
}