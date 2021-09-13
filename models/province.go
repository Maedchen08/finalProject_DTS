package models
import(
	"gorm.io/gorm"
)

type Province struct {
	gorm.Model
	Id uint `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:district_name;type:varchar(255);not null" json:"district_name"`
	District District `gorm:"foreignKey:province_id"`
	Regency Regency `gorm:"foreignKey:province_id"`
	User User `gorm:"foreignKey:province_id"`
	
}