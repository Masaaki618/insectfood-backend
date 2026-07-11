package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB は環境変数からDSNを組み立ててGORMのDB接続を生成する
func NewDB() (*gorm.DB, error) {
	tls := os.Getenv("DB_TLS")
	if tls == "" {
		tls = "false"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Asia%%2FTokyo&tls=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		tls,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("NewDB: %w", err)
	}
	return db, nil
}
