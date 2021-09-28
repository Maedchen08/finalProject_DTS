package models

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	Id       int   `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"column:username;type:varchar(255);not null;unique" json:"email"`
	Password string `gorm:"column:password;type:varchar(255);unique" json:"password"`

}

func (login *Login) TableName() string {
	return "login"
}
