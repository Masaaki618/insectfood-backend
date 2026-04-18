package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/models"
	"github.com/Masaaki618/insectfood-backend/internal/repositories"
)

const QuestionsPerCategory = 2 // パッケージ内のみ使用
var ErrInsufficientQuestions = errors.New("insufficient questions")

type questionService struct {
	repository repositories.IQuestionRepository
}

// NewQuestionService はIQuestionServiceを生成する
func NewQuestionService(repository repositories.IQuestionRepository) IQuestionService {
	return &questionService{repository: repository}
}

// GetQuestions はカテゴリ別にランダムで6問取得しDTOに詰め替えて返す
func (s *questionService) GetQuestions(ctx context.Context) ([]dtos.QuestionResponse, error) {
	var questionsRes []dtos.QuestionResponse
	visualQuestions, err := s.repository.GetRandomQuestionsByCategory(ctx, models.CategoryVisual, QuestionsPerCategory)
	if err != nil {
		return nil, fmt.Errorf("QuestionService.GetQuestions visual: %w", err)
	}
	if len(visualQuestions) < QuestionsPerCategory {
		return nil, ErrInsufficientQuestions
	}

	for _, visualQuestion := range visualQuestions {
		questionsRes = append(questionsRes, dtos.QuestionResponse{
			ID:       visualQuestion.ID,
			Body:     visualQuestion.Body,
			Category: string(visualQuestion.Category),
		})
	}

	physicalQuestions, err := s.repository.GetRandomQuestionsByCategory(ctx, models.CategoryPhysical, QuestionsPerCategory)
	if err != nil {
		return nil, fmt.Errorf("QuestionService.GetQuestions physical: %w", err)
	}
	if len(physicalQuestions) < QuestionsPerCategory {
		return nil, ErrInsufficientQuestions
	}

	for _, physicalQuestion := range physicalQuestions {
		questionsRes = append(questionsRes, dtos.QuestionResponse{
			ID:       physicalQuestion.ID,
			Body:     physicalQuestion.Body,
			Category: string(physicalQuestion.Category),
		})
	}
	mentalQuestions, err := s.repository.GetRandomQuestionsByCategory(ctx, models.CategoryMental, QuestionsPerCategory)
	if err != nil {
		return nil, fmt.Errorf("QuestionService.GetQuestions mental: %w", err)
	}
	if len(mentalQuestions) < QuestionsPerCategory {
		return nil, ErrInsufficientQuestions
	}

	for _, mentalQuestion := range mentalQuestions {
		questionsRes = append(questionsRes, dtos.QuestionResponse{
			ID:       mentalQuestion.ID,
			Body:     mentalQuestion.Body,
			Category: string(mentalQuestion.Category),
		})
	}

	return questionsRes, nil
}
