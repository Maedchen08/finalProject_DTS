package models

import (
	"gorm.io/gorm"
)

type Role struct { 
	gorm.Model
	Id       uint   `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	RoleName string `gorm:"column:role_name;type:varchar(10);not null" json:"role_name"`
	User User `gorm:"foreignKey:role_id"`
}
