package models

import "gorm.io/gorm"

type Agents struct {
	gorm.Model
	Id           int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string         `gorm:"column:agent_name;type:varchar(255);not null" json:"agent_name"`
	Address      string         `gorm:"column:address;type:varchar(255)" json:"alamat_agen"`
	ProvinceId   int            `gorm:"column:agent_province_id;not null" json:"agent_province_id"`
	RegencyId    int            `gorm:"column:agent_regency_id; not null" json:"agent_regency_id"`
	DistrictId   int            `gorm:"column:agent_district_id; not null" json:"agent_district_id"`
	NoWa         string         `gorm:"column:no_wa;type:varchar(30)" json:"no_wa"`
	Users        []Users        `gorm:"ForeignKey:agent_id" json:"-"`
	Transactions []Transactions `gorm:"ForeignKey:agents_id" json:"-"`
	LongtitudeAgent float64 `gorm:"column:longitude_agent" json:"longitude_agent"`
	LatitudeAgent float64 `gorm:"column:latitude_agent" json:"latitude_agent"`
	
}

func (agent *Agents) TableName() string {
	return "agent"
}


type ListRekomAgents struct {
	Id           int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string         `gorm:"column:agent_name;type:varchar(255);not null" json:"agent_name"`
	Address      string         `gorm:"column:address;type:varchar(255)" json:"alamat_agen"`
	ProvinceId   int            `gorm:"column:agent_province_id;not null" json:"agent_province_id"`
	RegencyId    int            `gorm:"column:agent_regency_id; not null" json:"agent_regency_id"`
	DistrictId   int            `gorm:"column:agent_district_id; not null" json:"agent_district_id"`
	NoWa         string         `gorm:"column:no_wa;type:varchar(30)" json:"no_wa"`
	LongtitudeAgent float64 `gorm:"column:longitude_agent" json:"longitude_agent"`
	LatitudeAgent float64 `gorm:"column:latitude_agent" json:"latitude_agent"`
	Distance float64 `json:"distance"`
}
