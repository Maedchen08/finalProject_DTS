package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint   `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:name_agen;type:varchar(255);not null" json:"name_agen"`
	Username string `gorm:"column:username;type:varchar(255);not null;unique" json:"email"`
	Password string `gorm:"column:password;type:varchar(255);unique" json:"password"`
	Address  string `gorm:"column:address;type:varchar(255)" json:"alamat_agen"`
	Province string `gorm:"column:province;type:varchar(30)" json:"province"`
	Regency  string `gorm:"column:regency;type:varchar(30)" json:"regency"`
	City     string `gorm:"column:city;type:varchar(30)" json:"city"`
	NoWa     string `gorm:"column:no_wa;;type:varchar(30)" json:"no_wa"`
	LoginAs  uint   `gorm:"column:login_as" json:"login_as"`
}
