## 背景 / 目的

各APIの実装が完了した後、バリデーションとエラーレスポンスが実装ごとにバラバラにならないよう
Gin のミドルウェアで統一されたエラーハンドリングを整備する。
**TDDで進め**、エラーパターンをテストで先に定義してから実装する。

- 依存: #5, #6, #7, #8
- ラベル: `backend`

---

## スコープ / 作業項目

TDD の **Red → Green → Refactor** サイクルで進める。

### 🔴 Red（テストを先に書く）

- `dto/error_dto.go` — 統一エラーレスポンス構造体の定義:
  ```go
  type ErrorResponse struct {
      Error ErrorDetail `json:"error"`
  }
  type ErrorDetail struct {
      Code    int    `json:"code"`
      Message string `json:"message"`
  }
  ```
- `controller/insect_controller_test.go`（HTTPテスト） — 以下を先に書く（この時点では失敗する）:
  - 存在しない ID → 404 かつ `{"error": {"code": 404, ...}}` 形式
  - 数字以外のパスパラメーター → 400 かつ統一形式
  - 診断スコア範囲外 → 400 かつ統一形式
  - エラーレスポンスにスタックトレースが含まれない

### 🟢 Green（テストが通る実装を書く）

- `internal/middleware/error_handler.go` — Gin カスタムエラーハンドラーの実装
- 各 Controller に散在しているエラーレスポンス処理をミドルウェアに集約
- パスパラメーター `:id` の型バリデーション（数字以外 → 400）
- `router/router.go` にミドルウェアを組み込む

### 🔵 Refactor（コードを整理する）

- エラーコード・メッセージの定数化（マジックストリング排除）

**参照設計書**:
- `docs/04_api.md`（エラーレスポンス形式・ステータスコード一覧）
- `docs/openapi.yml`（エラースキーマ定義）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `make test` でエラーハンドリングのテストがすべて通る（Green）
- [ ] 全エラーレスポンスが `{"error": {"code": N, "message": "..."}}` 形式に統一されている
- [ ] HTTP 400 / 404 / 500 / 503 それぞれで正しいステータスコードが返る
- [ ] エラーレスポンスにスタックトレース・DB接続文字列などの内部情報が含まれない
- [ ] パスパラメーター（`:id`）に文字列（例: `/api/v1/insects/abc`）が渡された場合に HTTP 400 を返す

---

## テスト観点

- ユニットテスト（ginkgo + httptest）:
  ```bash
  go test ./internal/controller/... -v
  ```
  - 各エラーパターンで期待したステータスコードとレスポンス形式を確認
- 手動確認（curl）:
  ```bash
  # 404
  curl -s http://localhost:8080/api/v1/insects/9999 | jq .

  # 400（型不正）
  curl -s http://localhost:8080/api/v1/insects/abc | jq .

  # 400（スコア範囲外）
  curl -s -X POST http://localhost:8080/api/v1/diagnosis \
    -H "Content-Type: application/json" \
    -d '{"scores": {"visual": 99}}' | jq .
  ```
