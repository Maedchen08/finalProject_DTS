package models
import(
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id uint `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:title;type:varchar(255);not null" json:"title" validate:"required"`
	Email string `gorm:"column:email;type:varchar(255);not null;unique" json:"email" validate:"required"`
	Password string `gorm:"column:password;type:varchar(255);unique" json:"password"`
	IsActive int `gorm:"column:is_active" json:"is_active" validate:"required"`
	RoleId uint `gorm:"column:role_id" json:"role_id" validate:"required"`
	AlamatLengkap string `gorm:"alamat_lengkap;type:varchar(255);unique" json:"alamat_lengkap"`
	ProvinceId uint `gorm:"province_id" json:"province_id"`
	DistrictId uint `gorm:"district_id" json:"district_id"`
	RegencyId uint `gorm:"regency_id" json:"regency"`
	NoWa string `gorm:"no_wa" json:"no_wa"`
	

}