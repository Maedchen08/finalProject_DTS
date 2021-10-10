package repositories

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/models/enum"
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
	DeleteTransaction(id int) error
	// AddAgent(transaction *models.Transactions) (int, error)
	GetByAgentId(id int) ([]models.Transactions,error)
	GetByCustomerId(id int) ([]models.Transactions,error)
	// SaveAndSearchAgent(transaction *models.Transactions)(int, error)

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
	err := tr.DB.Model(&models.Transactions{}).Where("id", t.Id).Update("status_transaction", enum.StatusTransaction.EnumIndex(1)).Error
	if err != nil {
		return err
	}
	return nil
	
}

//change status to reject
func (tr *TransactionRepo) ChangeReject(t *models.Transactions) error {
	err := tr.DB.Model(&models.Transactions{}).Where("id", t.Id).Update("status_transaction", enum.Canceled).Error
	if err != nil {
		return err
	}
	return nil
}

//change status to done
func (tr *TransactionRepo) ChangeDone(t *models.Transactions) error {
	err := tr.DB.Model(&models.Transactions{}).Where("id", t.Id).Update("status_transaction",enum.Done).Error
	if err != nil {
		return err
	}
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

//  delete transaction 
func (tr *TransactionRepo) DeleteTransaction(id int) error{
	var trans models.Transactions
	result := tr.DB.Where("id =?", id).Delete(&trans)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil

}

//cari agent
// func (tr *TransactionRepo) AddAgent(transaction *models.Transactions) (int,error){
// 	query := "INSERT INTO transaction "
// }

//Get transaksi by agent Id
func (tr *TransactionRepo) GetByAgentId(id int) ([]models.Transactions, error) {
	var transaction []models.Transactions
	transactionAgent := tr.DB.Where("agents_id = ?", id).Find(&transaction)
	return transaction, transactionAgent.Error
}

//Get transaksi by customer Id
func (tr *TransactionRepo) GetByCustomerId(id int) ([]models.Transactions, error) {
	var transaction []models.Transactions
	transactionAgent := tr.DB.Where("customers_id = ?", id).Find(&transaction)
	return transaction, transactionAgent.Error
}

//cari agent 
// func (tr *TransactionRepo) SaveAndSearchAgent(transaction *models.Transactions) (int, error){
// }

