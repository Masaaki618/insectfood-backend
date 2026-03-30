package repositories

import (
	"context"

	"github.com/Masaaki618/insectfood-backend/internal/models"
)

type IQuestionRepository interface {
	GetRandomQuestionsByCategory(ctx context.Context, category models.QuestionCategory, limit int) ([]models.Question, error)
}
