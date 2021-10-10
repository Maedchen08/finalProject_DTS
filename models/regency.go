package models

import (
	"gorm.io/gorm"
)

type Regency struct {
	gorm.Model
	ProvinceId int      `gorm:"column:province_regency_id" json:"province_regency_id"`
	Name       string   `gorm:"column:name;type:varchar(255)" json:"name"`
	District   District `gorm:"foreignKey:district_regency_id" json:"district_regency_id"`
	Agents  Agents  `gorm:"foreignKey:agent_regency_id" json:"agent_regency_id"`
	Transactions Transactions `gorm:"foreignKey:transaction_regency_id" json:"transaction_regency_id"`
}

func (regency *Regency) TableName() string {
	return "regency"
}
