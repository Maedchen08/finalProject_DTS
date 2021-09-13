package models


import(
	"gorm.io/gorm"
	"AntarJemput-Be-C/models/enum"

)

type Transaction struct {
	gorm.Model
	Id uint `gorm:"column:id; primaryKey;autoIncrement" json:"id"`
	Amount uint `gorm:"column:amount" json:"amount"`
	UserId uint `gorm:"column:user_id" json:"user_id"`
	Status enum.StatusTransaction `gorm:"column:status" json:"status"`
	TypeTransactionId uint `gorm:"column:type_transaction_id" json:"type_transaction_id"`


}