package models

import (
	"gorm.io/gorm"
)

type TypeServiceTransaction struct {
	gorm.Model
	Id                   int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TypeTransactionName  string         `gorm:"column:type_transaction_name;type:varchar(40);not null" json:"type_transaction_name"`
	ServiceTransactionId int            `gorm:"column:service_transaction_id" json:"service_transaction_id"`
	Transactions         []Transactions `gorm:"ForeignKey:type_transaction_id" json:"type_transaction_id"`
	
}

func (typeServiceTransaction *TypeServiceTransaction) TableName() string {
	return "type_service_transaction"
}
