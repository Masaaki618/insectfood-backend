## 背景 / 目的

Go + Gin の三層アーキテクチャ（Router → Controller → Service → Repository）の骨格を作り、
以降の各API実装がこの構造に乗れる状態にする。インターフェースによるDIと、
TDD で使う Ginkgo / gomock のセットアップもここで完結させる。

- 依存: #1
- ラベル: `backend`

---

## スコープ / 作業項目

- `go mod init` / 依存パッケージの追加:
  - Gin, sqlx, godotenv（アプリ本体）
  - **ginkgo, gomega**（テストフレームワーク）
  - **gomock, mockgen**（モック自動生成）
- `internal/` 配下のディレクトリ・空ファイルの作成:
  - `router/router.go`
  - `controller/` （空ファイル）
  - `service/` （インターフェースファイル含む）
  - `repository/` （インターフェースファイル含む）
  - `model/insect.go`, `model/question.go`, `model/radar_chart.go`
  - `dto/insect_dto.go`, `dto/question_dto.go`, `dto/diagnosis_dto.go`
  - `infrastructure/database/db.go`
  - `infrastructure/ai/claude_client.go`（空実装）
- `cmd/server/main.go` の作成（サーバー起動エントリーポイント）
- Ginkgo のブートストラップ（`ginkgo bootstrap` を各パッケージで実行）
- `Makefile` に以下を追加:
  - `make build` / `make run`
  - `make test`（ginkgo 実行）
  - `make mock`（mockgen で mock を再生成）

**参照設計書**: `docs/02_architecture.md`（フォルダ構成・三層アーキテクチャ設計）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `docs/02_architecture.md` のフォルダ構成通りにディレクトリ・ファイルが存在する
- [ ] 各層（Service / Repository）のインターフェースが別ファイルで定義されている
- [ ] `internal/infrastructure/database/db.go` で MySQL への接続確立・切断ができる
- [ ] `cmd/server/main.go` から Gin サーバーが起動し、Router が読み込まれる
- [ ] `make test` を実行するとテストスイートが（テストが0件でも）エラーなく起動する
- [ ] `go build ./...` と `go vet ./...` がエラーなく通る

---

## テスト観点

- 手動確認:
  - `go build ./...` でビルドエラーがないこと
  - `make test` で ginkgo が起動し「0 tests passed」のような出力が出ること
- 検証方法: `go vet ./...` を実行して警告がゼロであることを確認

---

## 要確認事項

- `go.mod` のモジュール名（例: `github.com/username/insectfood-backend`）を決める必要がある
- DB接続に `database/sql` を使うか `sqlx` を使うかを事前に決めておくと後続 Issue がスムーズ
