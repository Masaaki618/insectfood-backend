package models

import "gorm.io/gorm"

// Question は診断質問を表すモデル
type Question struct {
	gorm.Model
	Body     string `gorm:"column:body;not null"     json:"body"`
	Category string `gorm:"column:category;not null" json:"category"`
}
