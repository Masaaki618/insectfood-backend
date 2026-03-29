.PHONY: migrate-up migrate-down seed up down logs lint

# DB接続情報（.envから読み込み）
include .env
DB_URL=mysql://$(DB_USER):$(DB_PASSWORD)@tcp(127.0.0.1:$(DB_PORT))/$(DB_NAME)

# マイグレーション実行
migrate-up:
	migrate -path rdb/migrations -database "$(DB_URL)" up

# マイグレーション巻き戻し（1つ戻す）
migrate-down:
	migrate -path rdb/migrations -database "$(DB_URL)" down 1

# 初期データ投入
seed:
	docker compose exec -T db mysql -u$(DB_USER) -p$(DB_PASSWORD) $(DB_NAME) < rdb/seeds/seed.sql

# Docker操作
up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f app

# Lint実行
lint:
	golangci-lint run ./...
