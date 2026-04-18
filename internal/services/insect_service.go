package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
	"github.com/Masaaki618/insectfood-backend/internal/infrastructure/ai"
	"github.com/Masaaki618/insectfood-backend/internal/repositories"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type insectService struct {
	repository repositories.IInsectRepository
	claude     ai.IClaudeClient
}

// NewInsectService はIInsectServiceを生成する
func NewInsectService(repository repositories.IInsectRepository, claude ai.IClaudeClient) IInsectService {
	return &insectService{repository: repository, claude: claude}
}

// GetInsects は昆虫の一覧を取得しDTOに詰め替えて返す
func (s *insectService) GetInsects(ctx context.Context) ([]dtos.InsectResponse, error) {
	insects, err := s.repository.GetInsects(ctx)
	if err != nil {
		return nil, fmt.Errorf("InsectService.GetInsects: %w", err)
	}
	response := []dtos.InsectResponse{}
	for _, insect := range insects {
		response = append(response, dtos.InsectResponse{
			ID:           insect.ID,
			Name:         insect.Name,
			Difficulty:   insect.Difficulty,
			Introduction: insect.Introduction,
			Taste:        insect.Taste,
			Texture:      insect.Texture,
			InsectImg:    insect.InsectImg,
		})
	}

	return response, nil
}

// GetInsectByID は指定IDの昆虫詳細とレーダーチャートを取得しDTOに詰め替えて返す
func (s *insectService) GetInsectByID(ctx context.Context, insectID uint) (*dtos.InsectDetailResponse, error) {
	insect, err := s.repository.GetInsectByID(ctx, insectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("InsectService.GetInsectByID: %w", err)
	}

	var insectRes dtos.InsectResponse
	insectRes.ID = insect.ID
	insectRes.Name = insect.Name
	insectRes.Difficulty = insect.Difficulty
	insectRes.Introduction = insect.Introduction
	insectRes.Taste = insect.Taste
	insectRes.Texture = insect.Texture
	insectRes.InsectImg = insect.InsectImg

	var radarChartRes dtos.RadarChartResponse

	radarChart, err := s.repository.GetRadarChartByInsectID(ctx, insectID)
	if err != nil {
		return nil, fmt.Errorf("InsectService.GetRadarChartByInsectID: %w", err)
	}

	if radarChart != nil {
		radarChartRes.UmamiScore = radarChart.UmamiScore
		radarChartRes.BitterScore = radarChart.BitterScore
		radarChartRes.EguScore = radarChart.EguScore
		radarChartRes.FlavorScore = radarChart.FlavorScore
		radarChartRes.KimoScore = radarChart.KimoScore
	}

	aiComment, err := s.claude.GenerateInsectComment(ctx, insect)
	if err != nil {
		aiComment = fmt.Sprintf("まずは%sから始めてみましょう！", insect.Name)
	}
	var response dtos.InsectDetailResponse
	response.InsectResponse = insectRes
	response.RadarChart = radarChartRes
	response.AIComment = aiComment

	return &response, nil
}
