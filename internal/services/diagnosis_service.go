package services

import (
	"context"
	"fmt"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/infrastructure/ai"
	"github.com/Masaaki618/insectfood-backend/internal/repositories"
)

type diagnosisService struct {
	repository repositories.IInsectRepository
	claude     ai.IClaudeClient
}

// NewDiagnosisService はIDiagnosisServiceを生成する
func NewDiagnosisService(repository repositories.IInsectRepository, claude ai.IClaudeClient) IDiagnosisService {
	return &diagnosisService{repository: repository, claude: claude}
}

// Diagnose はスコアを受け取りAIがレコメンドした昆虫とコメントを返す（AI呼び出しは後で実装）
func (s *diagnosisService) Diagnose(ctx context.Context, req dtos.DiagnosisRequest) (*dtos.DiagnosisResponse, error) {
	insects, err := s.repository.GetInsects(ctx)
	if err != nil {
		return nil, fmt.Errorf("diagnosisService.Diagnose: %w", err)
	}

	insectID, aiComment, err := s.claude.GenerateDiagnosisResult(
		ctx,
		req.Scores.Visual,
		req.Scores.Physical,
		req.Scores.Mental,
		insects,
	)

	if err != nil {
		return nil, fmt.Errorf("diagnosisService.Diagnose: %w", err)
	}

	var insectRes dtos.InsectResponse
	for _, insect := range insects {
		if insect.ID == insectID {
			insectRes.ID = insectID
			insectRes.Name = insect.Name
			insectRes.Difficulty = insect.Difficulty
			insectRes.Introduction = insect.Introduction
			insectRes.Texture = insect.Texture
			insectRes.Taste = insect.Taste
			insectRes.InsectImg = insect.InsectImg
		}
	}

	var diagnosisRes dtos.DiagnosisResponse
	diagnosisRes.Insect = insectRes
	diagnosisRes.AIComment = aiComment
	return &diagnosisRes, nil
}
