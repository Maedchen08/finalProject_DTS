package models

import (
	"gorm.io/gorm"
)

type ServiceTransaction struct {
	gorm.Model
	Id       int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ServiceName string  `gorm:"column:service_name;type:varchar(30);not null" json:"service_name"`
	TypeServiceTransaction []TypeServiceTransaction `gorm:"ForeignKey:service_transaction_id" json:"-"`
	
}

func (serviceTransaction *ServiceTransaction) TableName() string {
	return "service_transaction"
}
