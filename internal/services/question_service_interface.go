package services

import (
	"context"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
)

// IQuestionService は質問に関するビジネスロジックを抽象化するインターフェース
type IQuestionService interface {
	// GetQuestions はカテゴリ別にランダムで6問取得する
	GetQuestions(ctx context.Context) ([]dtos.QuestionResponse, error)
}
