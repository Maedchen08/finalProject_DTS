package models

import (
	"gorm.io/gorm"
)

type District struct {
	gorm.Model
	ID         int    `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Name       string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	ProvinceId uint64 `gorm:"ForeignKey:province_id" json:"province_id"`
	// User User `gorm:"foreignKey:district_id"`

}
