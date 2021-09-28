package services

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/repositories"
)

type CustomerService struct {
	customerRepository repositories.CustomerRepoInterface
}

func NewCustomerService(customerRepository repositories.CustomerRepoInterface) *CustomerService {
	return &CustomerService{customerRepository: customerRepository}
}

type CustomerServiceInterface interface {
	Save(customer *models.Customer) (*models.Customer, error)
	GetById(id int) (models.Customer, error)
	GetAll() ([]models.Customer, error)
}

func (cs *CustomerService) Save(customer *models.Customer) (*models.Customer, error) {
	id, err := cs.customerRepository.Save(customer)
	if err != nil {
		return nil, err
	}
	customer.Id = id
	return customer, nil
}

func (cs *CustomerService) GetById(id int) (models.Customer, error) {
	customer, err := cs.customerRepository.GetById(id)
	if err != nil {
		return customer, err
	}
	return customer,nil
}

func (cs *CustomerService) GetAll() ([]models.Customer, error) {
	customer, err := cs.customerRepository.GetAll()
	return customer,err
}