## 背景 / 目的

Cloud Run にデプロイしたバックエンドが接続する本番 DB として、TiDB Cloud Starter クラスターを構築する。
マイグレーションと初期データ投入を行い、本番環境でAPIが正常動作する状態にする。

- 依存: #13
- ラベル: `infra`, `database`

---

## スコープ / 作業項目

- TiDB Cloud Starter クラスターの作成（Cloud Console）
- 本番用データベースの作成
- TiDB Cloud の管理画面（SQL Editor）から以下を順番に実行:
  1. `000001_create_insects_table.up.sql`
  2. `000002_create_radar_charts_table.up.sql`
  3. `000003_create_questions_table.up.sql`
  4. 初期データ（昆虫10種・質問12問）の INSERT SQL
- TiDB Cloud の接続情報を Cloud Run の環境変数に設定
- Cloud Run と TiDB Cloud 間の接続確認

**参照設計書**:
- `docs/03_database.md`（テーブル定義・初期データ・マイグレーション運用「本番環境」セクション）
- `docs/02_architecture.md`（TiDB Cloud 選定理由・初期コスト）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] TiDB Cloud Starter クラスターが作成されており、管理画面からアクセスできる
- [ ] 3テーブル（insects / radar_charts / questions）が本番DBに存在する
- [ ] 初期データ（昆虫10種・質問12問）が投入されており、SQL で件数確認できる
- [ ] Cloud Run から TiDB Cloud への疎通が確認できる（`GET /health` が 200 を返す）
- [ ] 本番DBの接続情報が Cloud Run の環境変数に設定されている（管理画面で確認可能）

---

## テスト観点

- 手動確認:
  ```bash
  # 本番環境でのヘルスチェック（DB疎通確認）
  curl https://<本番URL>/health
  # {"status": "ok"}

  # 昆虫一覧（初期データ確認）
  curl https://<本番URL>/api/v1/insects | jq '.data | length'
  # 10

  # 診断API（E2E動作確認）
  curl -X POST https://<本番URL>/api/v1/diagnosis \
    -H "Content-Type: application/json" \
    -d '{"scores": {"visual": 1, "physical": 1, "mental": 1}}' | jq .
  ```
- 検証方法: TiDB Cloud の SQL Editor で `SELECT COUNT(*) FROM insects;` を実行して件数を確認

---

## 要確認事項

- TiDB Cloud の無料枠（5GB）を超えないよう、本番データ量に注意（初期データは問題なし）
- TiDB Cloud の接続には SSL が必要な場合がある（接続文字列に `tls=true` 等のオプションが必要か確認）
- バックアップ設定: TiDB Cloud Starter の自動バックアップポリシーを確認しておく
