package ai

import (
	"context"

	"github.com/Masaaki618/insectfood-backend/internal/models"
)

type IClaudeClient interface {
	GenerateInsectComment(ctx context.Context, insect *models.Insect) (string, error)
	GenerateDiagnosisResult(ctx context.Context, visual, physical, mental uint8, insects []models.Insect) (uint, string, error)
}
