# 実装計画｜昆虫食初心者ガイド バックエンド

---

## フェーズ構成

| フェーズ | 目的 | 含むIssue |
|---|---|---|
| **Phase 1: Walking Skeleton** | 最小構成でAPIが1本通る状態を作る。開発環境・DB・ルーティングの骨格＋Ginkgo/gomockのセットアップを確立 | #1 〜 #4 |
| **Phase 2: コア機能実装（TDD）** | MVPの4エンドポイントをTDDで実装。各Issueで「テスト先行 → 実装 → リファクタリング」を完結させる | #5 〜 #9 |
| **Phase 3: 堅牢化** | エラーハンドリング・バリデーションをTDDで整備し、カバレッジ計測で品質を担保する | #10 〜 #12 |
| **Phase 4: デプロイ** | Cloud Run・TiDB Cloud本番環境を構築し、リリースできる状態にする | #13 〜 #14 |

---

## 依存関係マップ

```
#1（開発環境）
  └─→ #2（DBマイグレーション）
        └─→ #3（プロジェクト骨格 + Ginkgo/gomockセットアップ）
              ├─→ #9（Claude APIクライアント + mock生成）
              │     ├─→ #6（昆虫詳細API ※TDD）
              │     └─→ #8（診断API ※TDD）
              └─→ #4（ヘルスチェックAPI）
                    ├─→ #5（昆虫一覧API ※TDD）
                    │     └─→ #6（昆虫詳細API ※TDD）
                    ├─→ #7（質問取得API ※TDD）
                    │     └─→ #8（診断API ※TDD）
                    └─→ #10（バリデーション・エラーハンドリング ※TDD）
                          └─→ #11（カバレッジ計測・品質確認）
                                └─→ #12（Repository統合テスト）
                                      └─→ #13（Cloud Run構築）
                                            └─→ #14（TiDB Cloud本番構築）
```

> ※ 各実装 Issue（#5〜#10）は内部で Red → Green → Refactor を完結させる
> ※ #9（Claude APIクライアント）は #6 / #8 よりも先に着手し、mock を生成しておく

---

## Issueアウトライン表

### Issue #1: ローカル開発環境を構築する
**概要**: Docker Compose でGoアプリ・MySQLコンテナを立ち上げ、air によるホットリロードを有効にする。
**依存**: -
**ラベル**: `infra`, `backend`
**AC**:
- [x] `docker compose up` で Go アプリと MySQL が起動する
- [x] `http://localhost:8080` にアクセスして何らかのレスポンスが返る
- [x] air によるホットリロードが動作する（ファイル変更 → 自動再起動）
- [x] `.env.example` に必要な環境変数のキーを列挙している
- [x] `README.md` にローカル起動手順を記載している

### Issue #2: DBマイグレーション基盤を構築する
**概要**: golang-migrate を導入し、3テーブル（insects / radar_charts / questions）の作成マイグレーションを実装する。
**依存**: #1
**ラベル**: `backend`, `database`
**AC**:
- [x] `make migrate-up` で3テーブルが作成される
- [x] `make migrate-down` で3テーブルがロールバックされる
- [x] up / down ファイルがセットで `rdb/migrations/` に存在する
- [x] `docs/03_database.md` のカラム定義・コメントと一致している
- [x] 初期データ（昆虫10種・質問12問）を投入する seed ファイルまたは SQL が存在する

### Issue #3: Goプロジェクトの骨格を実装する（Ginkgo/gomockセットアップ含む）
**概要**: 三層アーキテクチャの空実装とインターフェースを作成し、TDDで使うGinkgo / gomockのセットアップも完結させる。
**依存**: #1
**ラベル**: `backend`
**AC**:
- [x] `docs/02_architecture.md` のフォルダ構成通りにディレクトリ・ファイルが存在する
- [x] 各層のインターフェースファイルが別ファイルで定義されている
- [x] `internal/infrastructure/database/db.go` でDB接続が確立できる
- [x] `main.go` から Router を経由して Gin が起動する
- [x] `make test` を実行するとテストスイートが（0件でも）エラーなく起動する
- [x] `go build` / `go vet` がエラーなく通る

### Issue #4: ヘルスチェックAPIを実装する
**概要**: `GET /health` でサーバーとDB接続の生存確認ができるエンドポイントを実装する。Walking Skeleton の完成形。
**依存**: #3
**ラベル**: `backend`
**AC**:
- [x] `GET /health` が `{"status": "ok"}` を返す
- [x] DB接続が切れている場合は `503` を返す
- [x] curl / Postman で動作確認できる

### Issue #5: 昆虫一覧API（GET /api/v1/insects）を実装する ※TDD
**概要**: TDDで進める最初のAPI実装。テストを先に書き（Red）、実装（Green）、整理（Refactor）の順で進める。以降のAPI実装の雛形となる。
**依存**: #4
**ラベル**: `backend`
**AC**:
- [x] `make test` で `insect_service_test.go` のテストがすべて通る（Green）
- [x] `GET /api/v1/insects` が配列形式 `[...]` で全昆虫を返す
- [x] `docs/04_api.md` のレスポンス形式と一致している
- [x] DBに昆虫データが0件のとき空配列 `[]` を返す（500にならない）
- [x] Controller / Service / Repository がインターフェースを介して依存している

### Issue #6: 昆虫詳細API（GET /api/v1/insects/:id）を実装する ※TDD
**概要**: TDDで実装。Claude APIのmock（#9で生成済み）を使い、AI連携を含むテストを先に書いてから実装する。
**依存**: #5, #9
**ラベル**: `backend`
**AC**:
- [ ] `make test` で `insect_service_test.go` の追加テストがすべて通る（Green）
- [ ] `GET /api/v1/insects/:id` が昆虫情報 + radar_chart + ai_comment を返す
- [ ] 存在しない ID を指定した場合に `404` を返す
- [ ] Claude API エラー時は 3回リトライし、失敗時はデフォルトコメントを返す
- [ ] `docs/04_api.md` のレスポンス形式と一致している

### Issue #7: 診断質問取得API（GET /api/v1/questions）を実装する ※TDD
**概要**: TDDで実装。カテゴリ別2問ずつ返すロジックをテストで先に定義してから実装する。
**依存**: #4
**ラベル**: `backend`
**AC**:
- [ ] `make test` で `question_service_test.go` のテストがすべて通る（Green）
- [ ] `GET /api/v1/questions` が6問（visual:2 / physical:2 / mental:2）を返す
- [ ] リクエストのたびに異なる問題の組み合わせが返る（ランダム性の確認）
- [ ] `docs/04_api.md` のレスポンス形式と一致している
- [ ] DBの各カテゴリに2問未満のデータしかない場合、適切なエラーを返す

### Issue #8: 診断API（POST /api/v1/diagnosis）を実装する ※TDD
**概要**: TDDで実装。診断ロジックのバグ防止が最重要のため、スコアの全パターン・AIエラー時の挙動をテストで先に定義してから実装する。
**依存**: #5, #9
**ラベル**: `backend`
**AC**:
- [ ] `make test` で `diagnosis_service_test.go` のテストがすべて通る（Green）
- [ ] `POST /api/v1/diagnosis` が insect + ai_comment を返す
- [ ] スコア（visual / physical / mental）が 0〜2 の範囲外のとき `400` を返す
- [ ] AI は必ず insects テーブルに登録済みの昆虫から選ぶ
- [ ] Claude API エラー時は 3回リトライし、失敗時はデフォルトレスポンスを返す

### Issue #9: Claude APIクライアントを実装する ※TDD
**概要**: TDDで実装し、mock（`mock_claude_client.go`）を生成する。後続の #6 / #8 がこのmockを使ってテストを書けるようにする。
**依存**: #3
**ラベル**: `backend`
**AC**:
- [ ] `ClaudeClient` インターフェースが定義されており、本実装とmockを切り替えられる
- [ ] `make mock` を実行すると `mock_claude_client.go` が自動生成される
- [ ] APIキーを環境変数から読み込んでいる（コードにハードコードしていない）
- [ ] リトライロジック（最大3回）が実装されている
- [ ] `.env.example` に `ANTHROPIC_API_KEY` キーを追加している

### Issue #10: バリデーション・エラーハンドリングを整備する ※TDD
**概要**: TDDで実装。エラーパターンをテストで先に定義し、Ginミドルウェアで統一されたエラーレスポンスを実装する。
**依存**: #5, #6, #7, #8
**ラベル**: `backend`
**AC**:
- [ ] `make test` でエラーハンドリングのテストがすべて通る（Green）
- [ ] 全エラーレスポンスが `{"error": {"code": N, "message": "..."}}` 形式で統一されている
- [ ] 400 / 404 / 500 / 503 それぞれで正しいステータスコードが返る
- [ ] エラーレスポンスにスタックトレースなどの内部情報が含まれない
- [ ] パスパラメーター（`:id`）に数字以外が渡された場合に 400 を返す

### Issue #11: テストカバレッジ計測・品質確認
**概要**: #5〜#10でTDDにより書かれたテストのカバレッジを計測し、抜け漏れがあれば補完する。Service層全体で70%以上を担保する。
**依存**: #8, #10
**ラベル**: `backend`, `test`
**AC**:
- [ ] `make coverage` でHTMLカバレッジレポートが出力される
- [ ] Service 層全体のカバレッジが 70% 以上
- [ ] DiagnosisService のスコア境界値（0/0/0 と 2/2/2）がテストされている
- [ ] Claude API の1回目失敗 → リトライ成功のケースがテストされている
- [ ] `make test` がエラーなく通る

### Issue #12: Repository層のユニットテストを実装する
**概要**: ginkgo を使い、Repository 層の DB 操作をインメモリ or テスト用 DB に対してテストする。
**依存**: #11
**ラベル**: `backend`, `test`
**AC**:
- [ ] `make test` で Repository テストが実行できる
- [ ] GetInsects / GetInsectByID / GetQuestions の正常系がテストされている
- [ ] 存在しない ID を指定したときの挙動がテストされている
- [ ] テスト用 DB は本番 DB と分離されている（環境変数で切り替え）

### Issue #13: Cloud Run に本番デプロイする
**概要**: Dockerfile を整備し、GitHub Actions または gcloud CLI で Cloud Run へデプロイするパイプラインを構築する。
**依存**: #12
**ラベル**: `infra`
**AC**:
- [ ] `docker build` でビルドが成功する（マルチステージビルド推奨）
- [ ] Cloud Run にデプロイしてヘルスチェック（`/health`）が通る
- [ ] 環境変数（DB接続情報・APIキー）が Cloud Run の Secret Manager または環境変数として設定されている
- [ ] 本番 URL が `https://your-api.run.app/api/v1` 形式で疎通できる
- [ ] `docs/02_architecture.md` に本番 URL を追記している

### Issue #14: TiDB Cloud 本番DBを構築する
**概要**: TiDB Cloud Starter のクラスターを作成し、本番用マイグレーションと初期データを投入する。
**依存**: #13
**ラベル**: `infra`, `database`
**AC**:
- [ ] TiDB Cloud Starter クラスターが作成されている
- [ ] 3テーブル（insects / radar_charts / questions）が本番DBに存在する
- [ ] 初期データ（昆虫10種・質問12問）が投入されている
- [ ] Cloud Run から TiDB Cloud への疎通が確認できる
- [ ] 本番DBの接続情報が Cloud Run の環境変数に設定されている

---

## 要確認事項

1. **初期データ（seed）の管理方法**: `make migrate-up` に含めるか、別の `make seed` コマンドにするか
2. **画像（insect_img）の管理**: 初期実装では `null` で進めるか、ダミー画像URLを入れるか
3. **テストカバレッジの閾値**: Service 層 70% を目標としているが、プロジェクト規模的に妥当か確認
4. **Claude API の請求管理**: 開発中の意図しない大量呼び出しを防ぐ上限設定を検討
5. **`DiagnosisResponse` のキー名**: `docs/04_api.md` は `"comment"` だが `openapi.yml` は `"ai_comment"` — どちらに統一するか要確認
