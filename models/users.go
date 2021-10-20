package models

import (
	"AntarJemput-Be-C/models/enum"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	Id         int       `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	RoleId     enum.Role `gorm:"column:role" json:"role"`
	CustomerId int       `gorm:"column:customer_id" json:"customer_id"`
	AgentId    int       `gorm:"column:agent_id" json:"agent_id"`
	Username   string    `gorm:"column:username;type:varchar(255);not null;unique" json:"username"`
	Password   string    `gorm:"column:password;type:text" json:"password"`
}

func (users *Users) TableName() string {
	return "users"
}
