## 背景 / 目的

このアプリのコア機能である診断機能を実装する。カテゴリ別スコアを受け取り、
Claude Haiku 4.5 が昆虫リストの中から最適な1匹をレコメンドしパーソナライズコメントを生成して返す。
**TDDで進める**。診断ロジックのバグを早期に検出するためテストが最も重要な Issue。

- 依存: #5, #9
- ラベル: `backend`

---

## スコープ / 作業項目

TDD の **Red → Green → Refactor** サイクルで進める。

### 🔴 Red（テストを先に書く）

- `dto/diagnosis_dto.go` — リクエスト/レスポンス DTO 定義
- `service/diagnosis_service_interface.go` — インターフェース定義
- `service/diagnosis_service_test.go` — 以下のテストケースを先に書く（この時点では失敗する）:
  - 正常系: 有効なスコア（例: visual=2, physical=1, mental=0）でレコメンドが返る
  - 正常系: 全スコア0（0/0/0）でもレコメンドが返る
  - 正常系: 全スコア最大（2/2/2）でもレコメンドが返る
  - 異常系: Claude API が3回失敗したときデフォルトレスポンス（コオロギパウダー）を返す
  - 異常系: Claude が返す JSON のパースに失敗したときデフォルトレスポンスを返す
  - 異常系: Repository（昆虫一覧取得）がエラーのときエラーを返す

### 🟢 Green（テストが通る実装を書く）

- `service/diagnosis_service.go` — 以下のビジネスロジックを実装:
  1. Repository から insects テーブルの全件取得
  2. スコアと昆虫リストを `docs/04_api.md` の「診断結果用プロンプト」に埋め込んで Claude API に送信
  3. Claude のレスポンス JSON をパースし `insect_id` で insects を特定
  4. エラー時はリトライ（最大3回）→ 失敗時はデフォルトレスポンス（コオロギパウダー固定）
- `controller/diagnosis_controller.go` — リクエストバリデーション・レスポンス整形
- `router/router.go` に `POST /api/v1/diagnosis` のルートを追加

### 🔵 Refactor（コードを整理する）

- プロンプト生成ロジックを別関数に切り出す（可読性向上）
- リトライロジックが claude_client と重複していれば整理

**参照設計書**:
- `docs/04_api.md`（診断APIのリクエスト/レスポンス形式・AIプロンプト設計「診断結果用」）
- `docs/01_requirements.md`（診断ロジック・AIレコメンドの流れ・エラー処理）
- `docs/openapi.yml`（`/diagnosis` スキーマ定義）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `make test` で `diagnosis_service_test.go` のテストがすべて通る（Green）
- [ ] `POST /api/v1/diagnosis` に `{"scores": {"visual": 2, "physical": 1, "mental": 0}}` を送ると insect + ai_comment を返す
- [ ] スコアが 0〜2 の範囲外（例: `"visual": 3`）のときに HTTP 400 を返す
- [ ] AI は必ず insects テーブルに登録済みの昆虫から `insect_id` を選ぶ（未登録の昆虫を返さない）
- [ ] Claude API エラー時は最大3回リトライし、失敗時はデフォルトレスポンス（コオロギパウダー + 固定コメント）を返す

---

## テスト観点

- ユニットテスト（ginkgo + gomock）:
  ```bash
  go test ./internal/service/... -v
  ```
  - 正常系: 各スコアパターンで `Diagnose()` が `DiagnosisResponse` を返す
  - 異常系: Claude mock が1回失敗 → リトライ → 成功するケース
  - 異常系: Claude mock が3回連続失敗 → デフォルトレスポンス
  - 異常系: Claude mock が不正な JSON を返す → デフォルトレスポンス
- 手動確認（curl）:
  ```bash
  curl -s -X POST http://localhost:8080/api/v1/diagnosis \
    -H "Content-Type: application/json" \
    -d '{"scores": {"visual": 2, "physical": 1, "mental": 0}}' | jq .
  ```

---

## 要確認事項

- `docs/04_api.md` の `"comment"` と `docs/openapi.yml` の `"ai_comment"` でキー名が不一致。
  どちらに統一するか確認が必要（`docs/implementation_plan.md` 要確認事項 #5 参照）
