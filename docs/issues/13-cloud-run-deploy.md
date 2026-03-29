## 背景 / 目的

テストが整備された後、Go バックエンドを Cloud Run に本番デプロイする。
コンテナイメージのビルドからデプロイまでの手順を整備し、本番 URL でAPIが疎通できる状態にする。

- 依存: #12
- ラベル: `infra`

---

## スコープ / 作業項目

- `Dockerfile` を本番用に整備（マルチステージビルドで軽量化）
- Cloud Run サービスの作成（`gcloud run deploy` または Cloud Console）
- 環境変数の設定:
  - `ANTHROPIC_API_KEY` — Cloud Run の環境変数 or Secret Manager
  - `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` — TiDB Cloud接続情報
- Cloud Run の設定:
  - リージョン: `asia-northeast1`（東京）推奨
  - ヘルスチェックパス: `/health`
  - 最小インスタンス: 0（コスト削減）
- `docs/02_architecture.md` に本番 URL を追記

**参照設計書**:
- `docs/02_architecture.md`（インフラ構成・Cloud Run 選定理由）
- `docs/04_api.md`（ベースURL形式: `https://your-api.run.app/api/v1`）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `docker build` でビルドが成功する（エラーなし）
- [ ] Cloud Run にデプロイした後、`GET https://<本番URL>/health` が `{"status": "ok"}` を返す
- [ ] 環境変数（DB接続情報・ANTHROPIC_API_KEY）が Cloud Run の設定に正しく反映されている（コードや Dockerfile にハードコードされていない）
- [ ] 本番 URL `https://<service>.run.app/api/v1` が疎通できる
- [ ] `docs/02_architecture.md` の `ベースURL` が実際の本番 URL に更新されている

---

## テスト観点

- 手動確認:
  ```bash
  # ヘルスチェック
  curl https://<本番URL>/health

  # 昆虫一覧（本番DBへの疎通確認）
  curl https://<本番URL>/api/v1/insects | jq .
  ```
- 検証方法: Cloud Run のログ（Cloud Console）でエラーが出ていないことを確認

---

## 要確認事項

- Google Cloud プロジェクト ID を確認する必要がある
- Cloud Run のリージョンを `asia-northeast1`（東京）に設定する（コスト・レイテンシの観点）
- `ANTHROPIC_API_KEY` は Secret Manager を使うか環境変数直接設定かを決める（Secret Manager推奨だがコスト発生に注意）
