package repositories

import (
	"AntarJemput-Be-C/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		DB: db,
	}
}

//represent the customer repository contract
type CustomerRepoInterface interface {
	Save(customer *models.Customer) (int, error)
	GetAll() ([]models.Customer, error)
	GetById(id int) (models.Customer, error)
}

func (cr *CustomerRepository) Save(customer *models.Customer) (int, error) {
	err := cr.DB.Create(customer).Error
	if err != nil {
		return customer.Id, err
	}
	return customer.Id, nil
}

func (cr *CustomerRepository) GetAll() ([]models.Customer, error) {
	var customer []models.Customer
	findCustomer := cr.DB.Find(&customer)
	return customer, findCustomer.Error
}

func (cr *CustomerRepository) GetById(id int) (models.Customer, error) {
	var customer models.Customer
	query := `SELECT*FROM customer WHERE id =?`
	err := cr.DB.Raw(query, id).Scan(&customer).Error

	if err != nil {
		return customer, err
	}
	if customer.Id == 0 {
		return customer, gorm.ErrRecordNotFound
	}
	return customer, nil
}
