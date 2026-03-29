package models

import "gorm.io/gorm"

// Insect は昆虫情報を表すモデル
type Insect struct {
	gorm.Model
	Name         string  `gorm:"not null" json:"name"`
	Difficulty   uint8   `gorm:"not null" json:"difficulty"`
	Introduction string  `gorm:"not null" json:"introduction"`
	Taste        string  `gorm:"not null" json:"taste"`
	Texture      string  `gorm:"not null" json:"texture"`
	InsectImg    *string `json:"insect_img"`
}
