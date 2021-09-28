package models

import "gorm.io/gorm"

type Agents struct {
	gorm.Model
	Id           int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string         `gorm:"column:agent_name;type:varchar(255);not null" json:"agent_name"`
	Address      string         `gorm:"column:address;type:varchar(255)" json:"alamat_agen"`
	Province     string         `gorm:"column:province;type:varchar(30)" json:"province"`
	Regency      string         `gorm:"column:regency;type:varchar(30)" json:"regency"`
	City         string         `gorm:"column:city;type:varchar(30)" json:"city"`
	NoWa         string         `gorm:"column:no_wa;type:varchar(30)" json:"no_wa"`
	Users        []Users        `gorm:"ForeignKey:agent_id" json:"agent_id"`
	Transactions []Transactions `gorm:"ForeignKey:agents_id" json:"agents_id"`
}

func (agent *Agents) TableName() string {
	return "agent"
}
