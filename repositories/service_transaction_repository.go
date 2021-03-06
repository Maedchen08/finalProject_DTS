package repositories

import (
	"AntarJemput-Be-C/models"

	"gorm.io/gorm"
)

type ServiceTransactionRepo struct {
	DB *gorm.DB
}

func NewServiceTransRepo(db *gorm.DB) *ServiceTransactionRepo {
	return &ServiceTransactionRepo{
		DB: db,
	}
}

type ServiceTransacInterface interface {
	Save(serviceTransaction *models.ServiceTransaction) (int, error)
	GetAll() ([]models.ServiceTransaction,error)
	GetById(id int) (models.ServiceTransaction, error)
}

func (str * ServiceTransactionRepo) Save(serviceTransaction *models.ServiceTransaction) (int, error) {
	err := str.DB.Create(serviceTransaction).Error
	if err != nil {
		return serviceTransaction.Id, err
	}
	return serviceTransaction.Id, nil
}

func (str *ServiceTransactionRepo) GetAll() ([]models.ServiceTransaction, error) {
	var serviceTransaction []models.ServiceTransaction
	findStr := str.DB.Find(&serviceTransaction)
	return serviceTransaction, findStr.Error	
}

 func (str *ServiceTransactionRepo) GetById(id int) (models.ServiceTransaction, error) {
	 var serviceTransaction models.ServiceTransaction
	 query := `SELECT * FROM service_transaction WHERE id =?`
	 err := str.DB.Raw(query, id).Scan(&serviceTransaction).Error

	 if err != nil {
		 return serviceTransaction, err
	 }
	 return serviceTransaction, nil
 }