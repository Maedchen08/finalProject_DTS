package models

import "gorm.io/gorm"

type District struct {
	gorm.Model
	RegencyId    int          `gorm:"column:district_regency_id;not null" json:"district_regency_id"`
	Name         string       `gorm:"column:name;type:varchar(255)" json:"name"`
	Agents       Agents       `gorm:"foreignKey:agent_district_id" json:"agent_district_id"`
	Transactions Transactions `gorm:"foreignKey:transaction_district_id" json:"transaction_district_id"`
}

func (district *District) TableName() string {
	return "district"
}
