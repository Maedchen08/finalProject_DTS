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
	ChangesRejected(t *models.Transactions) (models.Transactions, error)
	ChangesDone(t *models.Transactions) (models.Transactions, error)
	GetAll() ([]models.Transactions, error)
	GetById(id int) (models.Transactions, error)
	DeleteTransaction(id int) error
	//GetByIdAgen(id int) (models.Transactions, error)
	GetByAgentId(id int) ([]models.Transactions, error)
	GetByCustomerId(id int) ([]models.Transactions, error)
	RatingTransaction(t *models.Transactions) (models.Transactions, error)
	RatingAgent(id int) (models.AgentRating, error)
}

func (ts *TransactionService) Save(newTransaction *models.Transactions) (*models.Transactions, error) {
	id, err := ts.transaction.Save(newTransaction)
	if err != nil {
		return nil, err
	}
	newTransaction.Id = id
	return newTransaction, nil
}

//new test
// func (ts *TransactionService) GetByIdAgen(id int) (models.Transactions, error) {
// 	trans, err := ts.transaction.GetByIdAgen(id)
// 	if err != nil {
// 		return trans, err
// 	}
// 	return trans, nil
// }

//confirm
func (ts *TransactionService) ChangesConfirmed(t *models.Transactions) (trans models.Transactions, err error) {
	err = ts.transaction.ChangeConfirmed(t)
	if err != nil {
		return trans, err
	}
	trans, err = ts.transaction.GetById(t.Id)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

//reject
func (ts *TransactionService) ChangesRejected(t *models.Transactions) (trans models.Transactions, err error) {
	err = ts.transaction.ChangeReject(t)
	if err != nil {
		return trans, err
	}
	trans, err = ts.transaction.GetById(t.Id)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

//done
func (ts *TransactionService) ChangesDone(t *models.Transactions) (trans models.Transactions, err error) {
	err = ts.transaction.ChangeDone(t)
	if err != nil {
		return trans, err
	}
	trans, err = ts.transaction.GetById(t.Id)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

func (ts *TransactionService) GetAll() ([]models.Transactions, error) {
	transac, err := ts.transaction.GetAll()
	return transac, err
}

func (ts *TransactionService) GetById(id int) (models.Transactions, error) {
	trans, err := ts.transaction.GetById(id)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

//delete a transaction by Id
func (ts *TransactionService) DeleteTransaction(id int) error {
	err := ts.transaction.DeleteTransaction(id)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TransactionService) GetByAgentId(id int) ([]models.Transactions, error) {
	trans, err := ts.transaction.GetByAgentId(id)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

func (ts *TransactionService) GetByCustomerId(id int) ([]models.Transactions, error) {
	trans, err := ts.transaction.GetByCustomerId(id)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

//Rating Transaction
func (ts *TransactionService) RatingTransaction(t *models.Transactions) (trans models.Transactions, err error) {
	err = ts.transaction.RatingTransaction(t)
	if err != nil {
		return trans, err
	}
	trans, err = ts.transaction.GetById(t.Id)
	if err != nil {
		return trans, err
	}
	return trans, nil
}

func (ts *TransactionService) RatingAgent(id int) (models.AgentRating, error) {
	trans, err := ts.transaction.RatingAgent(id)
	if err != nil {
		return trans, err
	}
	return trans, nil
}