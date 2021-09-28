package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Id           int            `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	Name         string         `gorm:"column:customer_name;type:varchar(255);not null" json:"customer_name"`
	NoWa         string         `gorm:"column:no_wa;;type:varchar(30)" json:"no_wa"`
	Users        []Users        `gorm:"ForeignKey:customer_id" json:"customer_id"`
	Transactions []Transactions `gorm:"ForeignKey:customers_id" json:"customers_id"`
}

func (customer *Customer) TableName() string {
	return "customer"
}
