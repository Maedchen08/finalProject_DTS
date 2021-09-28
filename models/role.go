package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Id        int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleName  string         `gorm:"column:role_name;type:varchar(10);not null" json:"role_name"`
	Users     []Users        `gorm:"ForeignKey:role_id;not null" json:"role_id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (role *Role) TableName() string {
	return "role"
}
