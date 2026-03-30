package repositories

import (
	"context"

	"github.com/Masaaki618/insectfood-backend/internal/models"
)

type IInsectRepository interface {
	GetInsects(ctx context.Context) ([]models.Insect, error)
	GetInsectByID(ctx context.Context, insectID uint) (*models.Insect, error)
	GetRadarChartByInsectID(ctx context.Context, insectID uint) (*models.RadarChart, error)
}
