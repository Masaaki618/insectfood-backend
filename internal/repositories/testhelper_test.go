package repositories_test

import (
	"fmt"
	"os"

	"github.com/Masaaki618/insectfood-backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// newTestDB はテスト用DBに接続してマイグレーションを実行する
func newTestDB() (*gorm.DB, error) {
	host := getEnvOrDefault("TEST_DB_HOST", "127.0.0.1")
	port := getEnvOrDefault("TEST_DB_PORT", "3307")
	user := getEnvOrDefault("DB_USER", "insectfood")
	password := getEnvOrDefault("DB_PASSWORD", "")
	name := getEnvOrDefault("DB_NAME", "insectfood")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Asia%%2FTokyo",
		user, password, host, port, name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("newTestDB: %w", err)
	}

	// テーブルを自動作成
	if err := db.AutoMigrate(&models.Insect{}, &models.RadarChart{}, &models.Question{}); err != nil {
		return nil, fmt.Errorf("AutoMigrate: %w", err)
	}

	return db, nil
}

func getEnvOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
