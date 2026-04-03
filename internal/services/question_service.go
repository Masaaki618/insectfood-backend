package services

import (
	"context"
	"fmt"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/models"
	"github.com/Masaaki618/insectfood-backend/internal/repositories"
)

const questionsPerCategory = 2 // パッケージ内のみ使用

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
	visualQuestions, err := s.repository.GetRandomQuestionsByCategory(ctx, models.CategoryVisual, questionsPerCategory)
	if err != nil {
		return nil, fmt.Errorf("QuestionService.GetQuestions visual: %w", err)
	}

	for _, visualQuestion := range visualQuestions {
		questionsRes = append(questionsRes, dtos.QuestionResponse{
			ID:       visualQuestion.ID,
			Body:     visualQuestion.Body,
			Category: string(visualQuestion.Category),
		})
	}

	physicalQuestions, err := s.repository.GetRandomQuestionsByCategory(ctx, models.CategoryPhysical, questionsPerCategory)
	if err != nil {
		return nil, fmt.Errorf("QuestionService.GetQuestions physical: %w", err)
	}

	for _, physicalQuestion := range physicalQuestions {
		questionsRes = append(questionsRes, dtos.QuestionResponse{
			ID:       physicalQuestion.ID,
			Body:     physicalQuestion.Body,
			Category: string(physicalQuestion.Category),
		})
	}
	mentalQuestions, err := s.repository.GetRandomQuestionsByCategory(ctx, models.CategoryMental, questionsPerCategory)
	if err != nil {
		return nil, fmt.Errorf("QuestionService.GetQuestions mental: %w", err)
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
