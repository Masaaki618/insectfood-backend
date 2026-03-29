package models

import "gorm.io/gorm"

// Insect は昆虫情報を表すモデル
type Insect struct {
	gorm.Model
	Name         string  `gorm:"column:name;not null"         json:"name"`
	Difficulty   int8    `gorm:"column:difficulty;not null"   json:"difficulty"`
	Introduction string  `gorm:"column:introduction;not null" json:"introduction"`
	Taste        string  `gorm:"column:taste;not null"        json:"taste"`
	Texture      string  `gorm:"column:texture;not null"      json:"texture"`
	InsectImg    *string `gorm:"column:insect_img"            json:"insect_img"`
}
