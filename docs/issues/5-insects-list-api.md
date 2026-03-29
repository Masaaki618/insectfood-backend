## 背景 / 目的

昆虫一覧画面（`/insects`）が必要とする `GET /api/v1/insects` エンドポイントを実装する。
三層アーキテクチャの最初の実装例となるため、**TDDで進め**、以降のAPI実装の雛形とする。

- 依存: #4
- ラベル: `backend`

---

## スコープ / 作業項目

TDD の **Red → Green → Refactor** サイクルで進める。

### 🔴 Red（テストを先に書く）

- `model/insect.go` — insects テーブルのモデル定義
- `repository/insect_repository_interface.go` — `GetInsects()` インターフェース定義
- `service/insect_service_interface.go` — `GetInsects()` インターフェース定義
- `repository/mock/mock_insect_repository.go` — mockgen で自動生成
- `service/insect_service_test.go` — 以下のテストケースを先に書く（この時点では失敗する）:
  - 正常系: Repository が昆虫リストを返したとき、Service も同じリストを返す
  - 正常系: Repository が空リストを返したとき、Service も空リストを返す
  - 異常系: Repository がエラーを返したとき、Service もエラーを返す

### 🟢 Green（テストが通る実装を書く）

- `repository/insect_repository.go` — DB から全件取得する実装
- `service/insect_service.go` — Repository を呼び出す実装
- `controller/insect_controller.go` — リクエスト受付・レスポンス整形
- `dto/insect_dto.go` — レスポンス用 DTO 定義
- `router/router.go` — `GET /api/v1/insects` のルート登録

### 🔵 Refactor（コードを整理する）

- 重複・可読性の問題があれば修正（テストは引き続き通る状態を維持）

**参照設計書**:
- `docs/04_api.md`（レスポンス形式・フィールド定義）
- `docs/03_database.md`（insects テーブル定義）
- `docs/openapi.yml`（`/insects` スキーマ定義）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `make test` で `insect_service_test.go` のテストがすべて通る（Green）
- [ ] `GET /api/v1/insects` が `{"data": [...]}` 形式で全昆虫を返す
- [ ] レスポンスのフィールド（id / name / difficulty / introduction / taste / texture / insect_img）が `docs/04_api.md` と一致している
- [ ] DBに昆虫データが0件のとき `{"data": []}` を返す（500 エラーにならない）
- [ ] Controller / Service / Repository の責務が正しく分離されており、各層がインターフェースを介して依存している

---

## テスト観点

- ユニットテスト（ginkgo + gomock）:
  ```bash
  make test
  # または
  go test ./internal/service/... -v
  ```
  - 正常系: mock が2件返すとき `GetInsects()` が2件返す
  - 正常系: mock が空リストを返すとき `GetInsects()` が空リストを返す
  - 異常系: mock がエラーを返すとき `GetInsects()` がエラーを返す
- 手動確認（curl）:
  ```bash
  curl -s http://localhost:8080/api/v1/insects | jq .
  ```
  - `data` 配列に10件の昆虫データが返ること
