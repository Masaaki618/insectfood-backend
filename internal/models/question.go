package models

import "gorm.io/gorm"

// QuestionCategory は診断質問のカテゴリを表す型
type QuestionCategory string

const (
	CategoryVisual   QuestionCategory = "visual"
	CategoryPhysical QuestionCategory = "physical"
	CategoryMental   QuestionCategory = "mental"
)

// Question は診断質問を表すモデル
type Question struct {
	gorm.Model
	Body     string           `gorm:"not null" json:"body"`
	Category QuestionCategory `gorm:"not null" json:"category"`
}
