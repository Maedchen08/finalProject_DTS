package models

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	Id       int   `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"column:username;type:varchar(255);not null;unique" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);unique" json:"password"`
	UserId int `gorm:"column:user_id" json:"user_id"`

}

func (login *Login) TableName() string {
	return "login"
}
