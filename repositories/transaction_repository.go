package repositories

import (
	"AntarJemput-Be-C/models"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{
		DB: db,
	}
}

type TransactionRepoInterface interface{
	Save(transaction *models.Transactions) (int, error)
}

func (tr *TransactionRepo) Save(transaction *models.Transactions) (int, error) {
	err := tr.DB.Create(transaction).Error
	if err != nil{
		return transaction.Id, err
	}
	return transaction.Id, nil
}