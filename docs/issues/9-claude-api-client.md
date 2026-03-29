## 背景 / 目的

診断API（#8）と昆虫詳細API（#6）の両方で Claude Haiku 4.5 を呼び出すため、
共通の Claude API クライアントを先に整備する。
**TDD で進め**、mock を使って後続の Issue がすぐテストを書けるようにする。

- 依存: #3
- ラベル: `backend`

---

## スコープ / 作業項目

TDD の **Red → Green → Refactor** サイクルで進める。

### 🔴 Red（テストを先に書く）

- `infrastructure/ai/claude_client.go` に `ClaudeClient` インターフェースを定義:
  ```go
  type ClaudeClient interface {
      Call(ctx context.Context, prompt string) (string, error)
  }
  ```
- `infrastructure/ai/mock/mock_claude_client.go` — mockgen で自動生成
- `infrastructure/ai/claude_client_test.go` — 以下のテストを先に書く（この時点では失敗する）:
  - 正常系: 有効なプロンプトを渡すと文字列レスポンスが返る（実APIを叩く統合テスト。`-short` フラグで Skip 可能にする）
  - 異常系: 誤った API キーのときエラーを返す

### 🟢 Green（テストが通る実装を書く）

- `infrastructure/ai/claude_client.go` の本体実装:
  - Anthropic API への HTTP リクエスト（モデル: `claude-haiku-4-5`）
  - リトライロジック（最大3回、指数バックオフ推奨）
  - タイムアウト設定（推奨: 30秒）
  - エラー時のラッピング
- `.env.example` に `ANTHROPIC_API_KEY=` を追加

### 🔵 Refactor（コードを整理する）

- リトライロジックを汎用ヘルパー関数に切り出す（#8 の DiagnosisService とも共有可能にする）

**参照設計書**:
- `docs/01_requirements.md`（エラー処理: リトライ3回）
- `docs/04_api.md`（AIプロンプト設計・使用モデル）
- `docs/02_architecture.md`（`infrastructure/ai/` の配置）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `ClaudeClient` インターフェースが定義されており、本実装と mock を切り替えられる
- [ ] `make mock` を実行すると `mock_claude_client.go` が自動生成される
- [ ] `ANTHROPIC_API_KEY` を環境変数から読み込んでいる（コードに直接書いていない）
- [ ] リトライロジックが実装されており、API エラー時に最大3回まで再試行する
- [ ] `.env.example` に `ANTHROPIC_API_KEY=` のキーが追記されている

---

## テスト観点

- ユニットテスト（mock を使った動作確認）:
  ```bash
  go test ./internal/infrastructure/ai/... -v -short
  # -short フラグで実APIを叩くテストをスキップ
  ```
- 統合テスト（実際の API を叩く。手動で実施）:
  ```bash
  go test ./internal/infrastructure/ai/... -v
  # 実際に Anthropic API にリクエストが飛ぶ
  ```
- 検証方法: mock が後続の #6 / #8 テストで正しく差し替えられることを確認

---

## 要確認事項

- Anthropic の公式 Go SDK（`github.com/anthropics/anthropic-sdk-go`）を使うか、素の HTTP クライアントで実装するかを決める（SDK 推奨）
- Claude API の月次請求上限を Anthropic コンソールで設定することを推奨（開発中の意図しない大量呼び出し防止）
