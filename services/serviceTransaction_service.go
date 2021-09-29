package services

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/repositories"
)

type STransactionService struct {
	str repositories.ServiceTransacInterface
}

func NewSTransactionService(str repositories.ServiceTransacInterface) *STransactionService {
	return &STransactionService{str: str}
}

type STSInterface interface {
	Save(st *models.ServiceTransaction) (*models.ServiceTransaction, error)
	GetAll() ([]models.ServiceTransaction, error)
	GetById(id int) (models.ServiceTransaction, error)
}

func (sts *STransactionService) Save(serviceTransaction *models.ServiceTransaction) (*models.ServiceTransaction, error) {
	id, err := sts.str.Save(serviceTransaction)
	if err != nil {
		return nil, err
	}
	serviceTransaction.Id = id
	return serviceTransaction, nil
}

func (sts *STransactionService) GetAll() ([]models.ServiceTransaction, error) {
	serviceTransac,err:=sts.str.GetAll()
	return serviceTransac,err
}

func (sts *STransactionService) GetById(id int) (models.ServiceTransaction,error) {
	st,err :=sts.str.GetById(id)
	if err != nil {
		return st,err
	}
	return st, nil
}
