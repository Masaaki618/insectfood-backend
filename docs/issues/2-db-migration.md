## 背景 / 目的

アプリが使用する3テーブル（insects / radar_charts / questions）をコードで管理するために、
golang-migrate を導入してマイグレーション基盤を整備する。初期データ投入もあわせて行う。

- 依存: #1
- ラベル: `backend`, `database`

---

## スコープ / 作業項目

- golang-migrate の導入（`go.mod` への追加）
- 以下のマイグレーションファイルを `rdb/migrations/` に作成:
  - `000001_create_insects_table.up.sql` / `.down.sql`
  - `000002_create_radar_charts_table.up.sql` / `.down.sql`
  - `000003_create_questions_table.up.sql` / `.down.sql`
- 初期データ（昆虫10種・質問12問）を投入するSQLの作成
- `Makefile` に `make migrate-up` / `make migrate-down` コマンドを追加

**参照設計書**: `docs/03_database.md`（テーブル定義・初期データ一覧・マイグレーション運用）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `make migrate-up` を実行すると3テーブル（insects / radar_charts / questions）がDBに作成される
- [ ] `make migrate-down` を実行すると3テーブルがすべて削除（ロールバック）される
- [ ] `rdb/migrations/` に up / down のファイルがセットで6ファイル存在する
- [ ] テーブル定義が `docs/03_database.md` のカラム定義・型・コメントと完全に一致している
- [ ] 初期データ（昆虫10種・質問12問）が投入され、SELECT で件数を確認できる

---

## テスト観点

- 手動確認:
  - `make migrate-up` → 各テーブルの存在と件数を `SELECT COUNT(*)` で確認
  - `make migrate-down` → テーブルが消えることを確認
  - 再度 `make migrate-up` でべき等に動作することを確認
- 検証方法: `docker compose exec db mysql -u root -p` でMySQLに接続してテーブルを確認

---

## 要確認事項

- 初期データは migrate-up に含めるか（seed ファイル分離 or SQL埋め込み）、`make seed` として別コマンド化するかを要確認（`docs/implementation_plan.md` 要確認事項 #1 参照）
- `insect_img`（画像URL）は初期データに仮URL or NULL どちらで登録するか（`docs/implementation_plan.md` 要確認事項 #2 参照）
