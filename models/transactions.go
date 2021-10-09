package models

import (
	"AntarJemput-Be-C/models/enum"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	Id                int                    `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	TypeTransactionId int                    `gorm:"column:type_transaction_id" json:"type_transaction_id"`
	CustomerId        int                    `gorm:"column:customers_id" json:"customers_id"`
	AgentId           int                    `gorm:"column:agents_id" json:"agents_id"`
	Address           string                 `gorm:"column:address;type:varchar(255);not null" json:"address"`
	ProvinceId        int                    `gorm:"column:transaction_prov_id;not null" json:"transaction_prov_id"`
	RegencyId         int                    `gorm:"column:transaction_regency_id;not null" json:"transaction_regency_id"`
	DistrictId        int                    `gorm:"column:transaction_district_id;not null" json:"transaction_district_id"`
	Amount            int                    `gorm:"column:amount;not null" json:"amount" validate:"required,min=50000,max=10000000"`
	StatusTransaction enum.StatusTransaction `gorm:"column:status_transaction;not null;default:0" json:"status_transaction"`
	Rating            int                    `gorm:"column:rating" json:"rating" validate:"required,min=1,max=5"`
}

func (transaction *Transactions) TableName() string {
	return "transaction"
}
