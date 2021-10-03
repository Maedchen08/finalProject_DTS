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

type TransactionRepoInterface interface {
	Save(transaction *models.Transactions) (int, error)
	ChangeConfirmed(transaction *models.Transactions) error
	ChangeReject(transaction *models.Transactions) error
	ChangeDone(transaction *models.Transactions) error
	GetAll() ([]models.Transactions, error)
	GetById(id int) (models.Transactions, error)
}

func (tr *TransactionRepo) Save(transaction *models.Transactions) (int, error) {
	err := tr.DB.Create(transaction).Error
	if err != nil {
		return transaction.Id, err
	}
	return transaction.Id, nil
}

//change status to confirmed
func (tr *TransactionRepo) ChangeConfirmed(t *models.Transactions) error {
	err := tr.DB.Model(&models.Transactions{}).Where("id", t.Id).Update("status_transaction", "confirmed").Error
	if err != nil {
		return err
	}
	tr.DB.Save(&models.Transactions{})
	return nil
}

//change status to reject
func (tr *TransactionRepo) ChangeReject(t *models.Transactions) error {
	err := tr.DB.Model(&models.Transactions{}).Where("id", t.Id).Update("status_transaction", "reject").Error
	if err != nil {
		return err
	}
	tr.DB.Save(&models.Transactions{})
	return nil
}

//change status to confirmed
func (tr *TransactionRepo) ChangeDone(t *models.Transactions) error {
	err := tr.DB.Model(&models.Transactions{}).Where("id", t.Id).Update("status_transaction", "done").Error
	if err != nil {
		return err
	}
	tr.DB.Save(&models.Transactions{})
	return nil
}

//GetAll transaction
func (tr *TransactionRepo) GetAll() ([]models.Transactions, error) {
	var transaction []models.Transactions
	findTransaction := tr.DB.Find(&transaction)
	return transaction, findTransaction.Error
}

//Get transaksi by Id
func (tr *TransactionRepo) GetById(id int) (models.Transactions, error) {
	var transaction models.Transactions
	// var servis models.ServiceTransaction
	// query := `SELECT a.id, b.type_transaction_name, a.amount, a.customers_id, a.agents_id, a.address, a.province, a.regency, a.city, a.status_transaction FROM transaction a left JOIN type_service_transaction b ON a.type_transaction_id = b.id WHERE a.id=?`;
	query := `SELECT * FROM transaction WHERE id = ?`
	err := tr.DB.Raw(query, id).Scan(&transaction).Error

	if err != nil {
		return transaction, err
	}
	if transaction.Id == 0 {
		return transaction, gorm.ErrRecordNotFound
	}
	return transaction, nil
}
