package models

import "gorm.io/gorm"

// RadarChart は昆虫のレーダーチャート用スコアを表すモデル
type RadarChart struct {
	gorm.Model
	InsectID    uint   `gorm:"not null"            json:"insect_id"`
	Insect      Insect `gorm:"foreignKey:InsectID" json:"insect"`
	UmamiScore  uint8  `gorm:"not null"            json:"umami_score"`
	BitterScore uint8  `gorm:"not null"            json:"bitter_score"`
	EguScore    uint8  `gorm:"not null"            json:"egu_score"`
	FlavorScore uint8  `gorm:"not null"            json:"flavor_score"`
	KimoScore   uint8  `gorm:"not null"            json:"kimo_score"`
}
