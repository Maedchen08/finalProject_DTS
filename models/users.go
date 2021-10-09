package models

import (
	"AntarJemput-Be-C/models/enum"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	Id         int    `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	RoleId     enum.Role    `gorm:"column:role" json:"role"`
	CustomerId int    `gorm:"column:customer_id" json:"customer_id"`
	AgentId    int    `gorm:"column:agent_id" json:"agent_id"`
	Username   string `gorm:"column:username;type:varchar(10);not null;unique" json:"username" validate:"min=4"`
	Password   string `gorm:"column:password;unique;not null" json:"password" validate:"min=5"`

}
func (users *Users) TableName() string {
	return "users"
}


