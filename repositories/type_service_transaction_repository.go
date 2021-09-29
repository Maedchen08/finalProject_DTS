package repositories

import (
	"AntarJemput-Be-C/models"

	"gorm.io/gorm"
)

type TypeServiceRepo struct {
	DB *gorm.DB
}

func NewTypeServiceRepo(db *gorm.DB) *TypeServiceRepo{
	return &TypeServiceRepo{
		DB: db,
	}
}

type TypeServiceInterface interface {
	Save(typeService *models.TypeServiceTransaction)(int, error)
	GetAll() ([]models.TypeServiceTransaction, error)
	GetById(id int) (models.TypeServiceTransaction, error)
}

func(ts *TypeServiceRepo) Save(typeService *models.TypeServiceTransaction) (int,error){
	err := ts.DB.Create(typeService).Error
	if err != nil {
		return typeService.Id, err
	}
	return typeService.Id, nil
}

func (ts *TypeServiceRepo) GetAll() ([]models.TypeServiceTransaction,error) {
	var typeService []models.TypeServiceTransaction
	findts := ts.DB.Find(&typeService)
	return typeService, findts.Error
}

func (ts *TypeServiceRepo) GetById(id int) (models.TypeServiceTransaction,error) {
	var typeService models.TypeServiceTransaction
	query := `SELECT * FROM type_service_transaction WHERE id =?`
	err := ts.DB.Raw(query, id).Scan(&typeService).Error

	if err != nil {
		return typeService,err
	}

	return typeService, nil
}