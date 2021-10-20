package models

import "gorm.io/gorm"

type AgentSearch struct {
	gorm.Model
	TypeTransactionId int     `gorm:"column:type_transaction_id" json:"type_transaction_id"`
	CustomerId        int     `gorm:"column:customers_id" json:"customers_id"`
	Address           string  `gorm:"column:address;type:varchar(255);not null" json:"address"`
	ProvinceId        int     `gorm:"column:transaction_prov_id;not null" json:"transaction_prov_id"`
	RegencyId         int     `gorm:"column:transaction_regency_id;not null" json:"transaction_regency_id"`
	DistrictId        int     `gorm:"column:transaction_district_id;not null" json:"transaction_district_id"`
	Amount            int     `gorm:"column:amount;not null" json:"amount"`
	LongtitudeCust    float64 `gorm:"column:longitude_cust" json:"longitude_cust"`
	LatitudeCust      float64 `gorm:"column:latitude_cust" json:"latitude_cust"`
}
