package repositories

import (
	"context"
	"fmt"

	"github.com/Masaaki618/insectfood-backend/internal/models"
	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) GetRandomQuestionsByCategory(ctx context.Context, category models.QuestionCategory, limit int) ([]models.Question, error) {
	var questions []models.Question
	err := r.db.WithContext(ctx).Where("category = ?", category).Order("RAND()").Limit(limit).Find(&questions).Error
	if err != nil {
		return nil, fmt.Errorf("GetRandomQuestionsByCategory: %w", err)
	}

	return questions, nil
}
