## 背景 / 目的

Walking Skeleton の仕上げとして、サーバーとDB接続の生死確認ができるヘルスチェックエンドポイントを実装する。
Cloud Run のヘルスチェック設定にも使用するため、本番デプロイ前に必須の実装。

- 依存: #3
- ラベル: `backend`

---

## スコープ / 作業項目

- `GET /health` エンドポイントの実装
- DB接続確認ロジックの実装（`db.Ping()` 等）
- `router/router.go` への `/health` ルート登録
- DB接続失敗時の 503 レスポンス実装

**参照設計書**: `docs/02_architecture.md`（ルーティング構成）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `GET /health` が `{"status": "ok"}` と HTTP 200 を返す
- [ ] DB接続が切れている状態で `GET /health` を呼ぶと HTTP 503 を返す
- [ ] curl または Postman で動作確認できる（ローカル環境）

---

## テスト観点

- 手動確認:
  - DBが起動した状態で `curl http://localhost:8080/health` → `{"status": "ok"}` が返ること
  - DBコンテナを停止した状態で同じリクエストを送り、503 が返ること
- 検証方法:
  ```bash
  # 正常系
  curl -s http://localhost:8080/health
  # {"status":"ok"}

  # 異常系（DBコンテナ停止後）
  docker compose stop db
  curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/health
  # 503
  ```
