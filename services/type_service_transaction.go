package services

import (
	"AntarJemput-Be-C/models"
	"AntarJemput-Be-C/repositories"
)

type TSService struct {
	tsr repositories.TypeServiceInterface
}

func NewSTService(tsr repositories.TypeServiceInterface) *TSService {
	return &TSService{tsr: tsr}
}

type TSSInterface interface {
	Save(ts *models.TypeServiceTransaction) (*models.TypeServiceTransaction, error)
	GetById(id int) (models.TypeServiceTransaction, error)
	GetAll() ([]models.TypeServiceTransaction, error)
}

func (tss *TSService) Save(ts *models.TypeServiceTransaction) (*models.TypeServiceTransaction, error) {
	id, err := tss.tsr.Save(ts)
	if err != nil {
		return nil, err
	}
	ts.Id = id
	return ts, nil
}

func (tss *TSService) GetAll() ([]models.TypeServiceTransaction, error){
	typeService,err := tss.tsr.GetAll()
	return typeService,err
}

func (tss *TSService) GetById(id int) (models.TypeServiceTransaction, error){
	ts,err := tss.tsr.GetById(id)
	if err != nil {
		return ts,err
	}
	return ts,nil
}
