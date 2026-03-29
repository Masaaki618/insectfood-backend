# 昆虫食初心者ガイド バックエンド

Go + Gin で実装した REST API サーバー。

---

## 技術スタック

| 項目 | 内容 |
|---|---|
| 言語 | Go 1.25.1 |
| フレームワーク | Gin |
| DB | MySQL 8.0（TiDB Cloud Starter / ローカルは Docker） |
| AI | Claude Haiku 4.5 |
| デプロイ | Cloud Run |

---

## ローカル起動手順

### 1. 環境変数の設定

```bash
cp .env.example .env
```

`.env` を開いて以下を設定する：

| 変数名 | 説明 |
|---|---|
| `DB_PASSWORD` | MySQLのパスワード（任意の文字列）|
| `DB_ROOT_PASSWORD` | MySQLのrootパスワード（任意の文字列）|
| `ANTHROPIC_API_KEY` | Anthropic APIキー |

### 2. 起動

```bash
docker compose up
```

### 3. 動作確認

```bash
curl http://localhost:8080/
# → {"message":"ok"}
```

### 4. 停止

```bash
docker compose down
```

DBのデータも削除したい場合：

```bash
docker compose down -v
```

---

## ホットリロード

`air` によるホットリロードが有効。`.go` ファイルを変更すると自動でサーバーが再起動される。

---

## DB マイグレーション

```bash
make migrate-up    # マイグレーション実行
make migrate-down  # マイグレーション巻き戻し
```
