package models

import "gorm.io/gorm"

// RadarChart は昆虫のレーダーチャート用スコアを表すモデル
type RadarChart struct {
	gorm.Model
	InsectID    uint   `gorm:"column:insect_id;not null"    json:"insect_id"`
	Insect      Insect `gorm:"foreignKey:InsectID"          json:"insect"`
	UmamiScore  int8   `gorm:"column:umami_score;not null"  json:"umami_score"`
	BitterScore int8   `gorm:"column:bitter_score;not null" json:"bitter_score"`
	EguScore    int8   `gorm:"column:egu_score;not null"    json:"egu_score"`
	FlavorScore int8   `gorm:"column:flavor_score;not null" json:"flavor_score"`
	KimoScore   int8   `gorm:"column:kimo_score;not null"   json:"kimo_score"`
}
