package services

import (
	"context"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
)

// IDiagnosisService は診断に関するビジネスロジックを抽象化するインターフェース
type IDiagnosisService interface {
	// Diagnose はスコアを受け取りAIがレコメンドした昆虫とコメントを返す
	Diagnose(ctx context.Context, req dtos.DiagnosisRequest) (*dtos.DiagnosisResponse, error)
}
