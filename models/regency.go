package models
import (
	"gorm.io/gorm"
)

type Regency struct {
	gorm.Model
	Id uint `gorm:"id;autoincrement;primarykey" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	ProvinceId uint64 `gorm:"ForeignKey:province_id" json:"province_id"`
	// User User `gorm:"foreignKey:regency_id"`

	
}