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
	Amount            int                    `gorm:"column:amount;not null" json:"amount"`
	StatusTransaction enum.StatusTransaction `gorm:"column:status_transaction;not null;default:0" json:"status_transaction"`
	Rating            int                    `gorm:"column:rating" json:"rating"`
	RatingComment     string           		 `gorm:"column:rating_comment;type:varchar(255);null" json:"rating_comment"`
	LongtitudeCust float64 `gorm:"column:longitude_cust" json:"longitude_cust"`
	LatitudeCust float64 `gorm:"column:latitude_cust" json:"latitude_cust"`
}

func (transaction *Transactions) TableName() string {
	return "transaction"
}
