package models

import (
	"gorm.io/gorm"
)

type Province struct {
	gorm.Model
	Name         string       `gorm:"column:name;type:varchar(255)" json:"name"`
	Regency      Regency      `gorm:"foreignKey:province_regency_id" json:"province_regency_id"`
	Agents       Agents       `gorm:"foreignKey:agent_province_id" json:"agent_province_id"`
	Transactions Transactions `gorm:"foreignKey:transaction_prov_id" json:"transaction_prov_id"`
}

func (province *Province) TableName() string {
	return "province"
}
