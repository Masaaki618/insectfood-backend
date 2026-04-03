package services

import (
	"context"

	"github.com/Masaaki618/insectfood-backend/internal/dtos"
)

// IInsectService は昆虫に関するビジネスロジックを抽象化するインターフェース
type IInsectService interface {
	// GetInsects は昆虫の一覧を取得する
	GetInsects(ctx context.Context) ([]dtos.InsectResponse, error)
	// GetInsectByID は指定IDの昆虫詳細とレーダーチャートを取得する
	GetInsectByID(ctx context.Context, insectID uint) (*dtos.InsectDetailResponse, error)
}
