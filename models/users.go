package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id         int `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	RoleId     int `gorm:"column:role_id" json:"role_id"`
	CustomerId int `gorm:"column:customer_id" json:"customer_id"`
	AgentId    int `gorm:"column:agent_id" json:"agent_id"`
	Login []Login `gorm:"ForeignKey:user_id" json:"user_id"`

}

func (users *Users) TableName() string {
	return "user"
}
