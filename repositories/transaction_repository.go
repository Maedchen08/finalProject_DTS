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
	ChangeConfirmed(transaction *models.Transactions) error
	GetAll() ([]models.Transactions, error)
}

func (tr *TransactionRepo) Save(transaction *models.Transactions) (int, error) {
	err := tr.DB.Create(transaction).Error
	if err != nil{
		return transaction.Id, err
	}
	return transaction.Id, nil
}

func (tr *TransactionRepo) ChangeConfirmed(t *models.Transactions) error {
	err := tr.DB.Model(&models.Transactions{}).Where("id",t.Id).Update("status_transaction","confirmed").Error
	if err != nil {
		return err
	}
	tr.DB.Save(&models.Transactions{})
	return  nil
}

func (tr *TransactionRepo) GetAll() ([]models.Transactions, error) {
	var transaction []models.Transactions
	findTransaction := tr.DB.Find(&transaction)
	return transaction, findTransaction.Error
}




