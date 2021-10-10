package models

import (
	"gorm.io/gorm"
)

type AgentRating struct {
	gorm.Model
	AgentId   int `gorm:"column:agents_id" json:"agents_id"`
	Total     int `gorm:"column:total; primaryKey;autoIncrement" json:"total"`
	AvgRating float32 `gorm:"column:avg_rating" json:"avg_rating"`
}

func (agen_rating *AgentRating) TableName() string {
	return "agen_rating"
}
