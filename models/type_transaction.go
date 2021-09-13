package models

import (
	"gorm.io/gorm"
)

type TypeTransaction struct { 
	gorm.Model
	Id       uint   `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	TypeName string `gorm:"column:rype_name;type:varchar(10);not null" json:"type_name"`
	Transaction Transaction `gorm:"foreignKey:type_transaction_id"`
	
}