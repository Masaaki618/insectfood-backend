package services

import (
	"context"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/repositories"
)

type diagnosisService struct {
	repository repositories.IInsectRepository
}

// NewDiagnosisService はIDiagnosisServiceを生成する
func NewDiagnosisService(repository repositories.IInsectRepository) IDiagnosisService {
	return &diagnosisService{repository: repository}
}

// Diagnose はスコアを受け取りAIがレコメンドした昆虫とコメントを返す（AI呼び出しは後で実装）
func (d *diagnosisService) Diagnose(ctx context.Context, req dtos.DiagnosisRequest) (*dtos.DiagnosisResponse, error) {
	var diagnosisRes dtos.DiagnosisResponse
	return &diagnosisRes, nil
}
