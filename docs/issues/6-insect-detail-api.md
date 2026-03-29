## 背景 / 目的

昆虫詳細画面（`/insects/:id`）が必要とする `GET /api/v1/insects/:id` エンドポイントを実装する。
insects と radar_charts の JOIN 取得に加え、Claude Haiku 4.5 による食レポコメントを付与する。
**TDDで進める**。

- 依存: #5, #9
- ラベル: `backend`

---

## スコープ / 作業項目

TDD の **Red → Green → Refactor** サイクルで進める。

### 🔴 Red（テストを先に書く）

- `model/radar_chart.go` — radar_charts テーブルのモデル定義
- `repository/insect_repository_interface.go` に `GetInsectByID()` を追加
- `service/insect_service_interface.go` に `GetInsectByID()` を追加
- `infrastructure/ai/mock/mock_claude_client.go` — mockgen で自動生成（#9 で作成済みなら流用）
- `service/insect_service_test.go` に以下を追加（この時点では失敗する）:
  - 正常系: 存在する ID を渡すと昆虫情報 + radar_chart + ai_comment を返す
  - 異常系: 存在しない ID を渡すと `sql.ErrNoRows` 系エラーを返す
  - 異常系: Claude API が3回失敗したときデフォルトコメントを返す

### 🟢 Green（テストが通る実装を書く）

- `repository/insect_repository.go` に `GetInsectByID()` を追加（insects + radar_charts JOIN）
- `service/insect_service.go` に `GetInsectByID()` を追加（Claude API 呼び出し含む）
- `controller/insect_controller.go` に `GET /api/v1/insects/:id` ハンドラーを追加
- `dto/insect_dto.go` に詳細レスポンス用 DTO を追加（ai_comment + radar_chart フィールド）
- `router/router.go` に `GET /api/v1/insects/:id` のルートを追加

### 🔵 Refactor（コードを整理する）

- `GetInsects()` と `GetInsectByID()` で重複している部分があれば共通化

**参照設計書**:
- `docs/04_api.md`（昆虫詳細APIのレスポンス形式・AIプロンプト設計「昆虫詳細用」）
- `docs/03_database.md`（radar_charts テーブル定義）
- `docs/openapi.yml`（`/insects/{id}` スキーマ定義）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `make test` で `insect_service_test.go` の追加テストがすべて通る（Green）
- [ ] `GET /api/v1/insects/:id` が昆虫情報 + radar_chart スコア + ai_comment を返す
- [ ] 存在しない ID（例: `/api/v1/insects/9999`）を指定した場合に HTTP 404 を返す
- [ ] Claude API エラー時は最大3回リトライし、失敗時はデフォルトコメントを返す
- [ ] レスポンス形式が `docs/04_api.md` および `docs/openapi.yml` の定義と一致している

---

## テスト観点

- ユニットテスト（ginkgo + gomock）:
  ```bash
  go test ./internal/service/... -v
  ```
  - 正常系: mock が昆虫データを返したとき ai_comment 付きのレスポンスを返す
  - 異常系: mock が `ErrNoRows` を返したとき Service がエラーを返す
  - 異常系: Claude mock が3回エラーを返したときデフォルトコメントになる
- 手動確認（curl）:
  ```bash
  curl -s http://localhost:8080/api/v1/insects/1 | jq .
  curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/api/v1/insects/9999
  # 404
  ```

---

## 要確認事項

- `docs/04_api.md` の `"comment"` と `docs/openapi.yml` の `"ai_comment"` でキー名が不一致。
  どちらに統一するか確認が必要（`docs/implementation_plan.md` 要確認事項 #5 参照）
