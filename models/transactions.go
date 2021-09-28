package models

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	Id                int    `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	TypeTransactionId int    `gorm:"column:type_transaction_id" json:"type_transaction_id"`
	CustomerId        int    `gorm:"column:customers_id" json:"customers_id"`
	AgentId           int    `gorm:"column:agents_id" json:"agents_id"`
	Address           string `gorm:"column:address;type:varchar(255);not null" json:"address"`
	Province          string `gorm:"column:province;type:varchar(20);not null" json:"province"`
	Regency           string `gorm:"column:regency;type:varchar(30)" json:"regency"`
	City              string `gorm:"column:city;type:varchar(30)" json:"city"`
	Amount            int    `gorm:"column:amount;not null" json:"amount"`
	StatusTransaction string `gorm:"column:status_transaction;type:varchar(20);not null" json:"status_transaction"`
	Rating            int    `gorm:"column:rating" json:"rating"`
}

func (transaction *Transactions) TableName() string {
	return "transaction"
}
