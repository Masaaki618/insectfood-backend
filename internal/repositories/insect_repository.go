package repositories

import (
	"context"
	"fmt"

	"github.com/Masaaki618/insectfood-backend/internal/models"
	"gorm.io/gorm"
)

type insectRepository struct {
	db *gorm.DB
}

func NewInsectRepository(db *gorm.DB) IInsectRepository {
	return &insectRepository{db: db}
}

func (r *insectRepository) GetInsects(ctx context.Context) ([]models.Insect, error) {
	var insects []models.Insect
	err := r.db.WithContext(ctx).Find(&insects).Error
	if err != nil {
		return nil, fmt.Errorf("GetInsects: %w", err)
	}
	return insects, nil
}

func (r *insectRepository) GetInsectByID(ctx context.Context, insectID uint) (*models.Insect, error) {
	var insect models.Insect
	err := r.db.WithContext(ctx).First(&insect, insectID).Error
	if err != nil {
		return nil, fmt.Errorf("GetInsectByID: %w", err) // ✓
	}
	return &insect, nil
}

func (r *insectRepository) GetRadarChartByInsectID(ctx context.Context, insectID uint) (*models.RadarChart, error) {
	var radarChart models.RadarChart
	err := r.db.WithContext(ctx).Where("insect_id = ?", insectID).First(&radarChart).Error
	if err != nil {
		return nil, fmt.Errorf("GetRadarChartByInsectID: %w", err)
	}
	return &radarChart, nil
}
