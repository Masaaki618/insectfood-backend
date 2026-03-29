## 背景 / 目的

開発を始めるにあたり、チーム全員が同じ環境でGoアプリとMySQLを起動できる基盤が必要。
Docker Compose + air によるホットリロード環境を整備し、「クローンしてすぐ動く」状態を作る。

- 依存: -
- ラベル: `infra`, `backend`

---

## スコープ / 作業項目

- `docker-compose.yml` の作成（Go アプリコンテナ + MySQL コンテナ）
- `Dockerfile` の作成（開発用: air ベース）
- `.air.toml` の設定（ホットリロード設定）
- `.env.example` の作成（必要な環境変数キーを列挙）
- `Makefile` に `make up` / `make down` コマンドを追加
- `README.md` にローカル起動手順を記載

**参照設計書**: `docs/02_architecture.md`（技術スタック・フォルダ構成）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `docker compose up` で Go アプリと MySQL が起動する
- [ ] `http://localhost:8080` にアクセスして何らかのレスポンスが返る（200 or 404 どちらでも可）
- [ ] air によるホットリロードが動作する（`.go` ファイルを変更するとコンテナが自動再起動する）
- [ ] `.env.example` に `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` などの必要なキーを列挙している
- [ ] `README.md` に「前提条件（Docker のインストール）」と「起動手順」が日本語で記載されている

---

## テスト観点

- 手動確認:
  - `docker compose up` を実行してエラーが出ないこと
  - `http://localhost:8080` に curl でリクエストして応答が返ること
  - Go ファイルを書き換えてコンテナが自動再起動することを確認
- 検証方法: ターミナルで `curl http://localhost:8080` を実行して応答を確認

---

## 要確認事項

- MySQL のバージョンを何にするか（TiDB Cloud が MySQL 5.7 互換のため、ローカルも 5.7 系が望ましい）
- `.env` ファイルは `.gitignore` に追加済みか確認が必要
