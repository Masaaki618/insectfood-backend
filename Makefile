.PHONY: migrate-up migrate-down seed up up-mock down logs lint test test-repo mock coverage

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
	docker compose exec -T db mysql -u$(DB_USER) -p$(DB_PASSWORD) --default-character-set=utf8mb4 $(DB_NAME) < rdb/seeds/seed.sql

# Docker操作
up:
	docker compose up

# モックAIモードで起動（Claude APIを叩かない）
up-mock:
	USE_MOCK_AI=true docker compose up

down:
	docker compose down

logs:
	docker compose logs -f app

# テスト実行（Service層・Controller層）
test:
	ginkgo -r ./internal/services/... ./internal/controllers/...

# Repository層のテスト実行（テスト用DBが必要）
test-repo:
	docker compose up -d db-test
	TEST_DB_HOST=127.0.0.1 TEST_DB_PORT=3307 DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_NAME=$(DB_NAME) ginkgo -r ./internal/repositories/...

# モック自動生成
mock:
	mockgen -source=internal/repositories/insect_repository_interface.go -destination=internal/repositories/mock/mock_insect_repository.go -package=mock
	mockgen -source=internal/repositories/question_repository_interface.go -destination=internal/repositories/mock/mock_question_repository.go -package=mock
	mockgen -source=internal/infrastructure/ai/claude_client_interface.go -destination=internal/infrastructure/ai/mock/mock_claude_client.go -package=mock

# カバレッジ計測（HTMLレポート出力）
coverage:
	go test -coverprofile=coverage.out ./internal/services/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "カバレッジレポート: coverage.html"

# Lint実行
lint:
	golangci-lint run ./...
