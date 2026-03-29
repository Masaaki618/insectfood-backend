## 背景 / 目的

Repository 層の DB 操作が正しく動作することを自動テストで保証する。
テスト専用のDBを用意して実際にクエリを発行するインテグレーションテストを実装し、
本番と同じ MySQL 互換環境での動作を確認する。

- 依存: #11
- ラベル: `backend`, `test`

---

## スコープ / 作業項目

- テスト用 DB の設定（`docker-compose.test.yml` または環境変数で切り替え）
- 以下の Repository テストファイルを作成:
  - `repository/insect_repository_test.go`
    - `GetInsects()` — 全件取得
    - `GetInsectByID()` — 正常系・存在しない ID
  - `repository/question_repository_test.go`
    - `GetRandomQuestions()` — カテゴリ別2問ずつ取得
- テスト実行前にテスト用DBのセットアップ（マイグレーション + テストデータ投入）を行う仕組みの実装
- `Makefile` に `make test-repo` コマンドを追加（または `make test` に統合）

**参照設計書**:
- `docs/03_database.md`（テーブル定義・初期データ）
- `docs/02_architecture.md`（テスト方針）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `make test` で Repository テストが実行できる
- [ ] `GetInsects()` / `GetInsectByID()` / `GetRandomQuestions()` の正常系がテストされている
- [ ] 存在しない ID を指定した `GetInsectByID()` がエラー（`sql.ErrNoRows` 等）を返すことをテストしている
- [ ] テスト用 DB は本番 DB と分離されており、環境変数（例: `TEST_DB_NAME`）で切り替えられる

---

## テスト観点

- インテグレーションテスト（テスト用MySQLに対して実際にクエリを発行）:
  - 正常系: テストデータを挿入して `GetInsects()` が全件返すことを確認
  - 正常系: 特定 ID を挿入して `GetInsectByID()` が正しいレコードを返すことを確認
  - 異常系: 存在しない ID に対して適切なエラーが返ることを確認
  - 正常系: `GetRandomQuestions()` がカテゴリ別2問ずつ返すことを確認
- 検証方法:
  ```bash
  # テスト用DBを起動してテスト実行
  docker compose -f docker-compose.test.yml up -d
  go test ./internal/repository/... -v
  ```

---

## 要確認事項

- テスト用 DB のデータ初期化戦略: テストケースごとにリセットするか（`BeforeEach` でtruncate）、テスト開始時に1回だけ初期化するか
- CI（GitHub Actions）でテストを自動実行するか（Phase 4 のデプロイ整備とセットで検討）
