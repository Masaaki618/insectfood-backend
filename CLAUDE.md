# CLAUDE.md｜昆虫食初心者ガイド バックエンド

Claude Code がこのプロジェクトで作業する際に守るべきルールと設定をまとめたファイルです。

---

## 1. 基本姿勢

- **必ず日本語で回答する**（コードのコメントも日本語）
- 「なぜそうするか」の理由も一緒に説明する
- 不明点があれば実装前に必ず確認する

---

## 2. プロジェクト概要

| 項目     | 内容                                          |
|--------|---------------------------------------------|
| アプリ名   | 昆虫食初心者ガイド                                   |
| バックエンド | Go + Gin（Cloud Run へデプロイ）                   |
| DB     | MySQL 互換（TiDB Cloud Starter / ローカルは Docker） |
| AI API | Claude Haiku 4.5（昆虫レコメンド・コメント生成）            |
| 設計書    | `docs/` 配下を参照                               |

### アーキテクチャ（三層構造）

```
Router → Controller → Service → Repository → DB
                          ↕
                   Claude Haiku 4.5
```

各層はインターフェースを介して依存する（GoのDuck Typing活用）。

### フォルダ構成（主要部分）

```
internal/
├── routers/        # ルーティング定義
├── controllers/    # リクエスト/レスポンス処理
├── services/       # ビジネスロジック + Claude API呼び出し
├── repositories/   # DBアクセス
├── models/         # DBモデル定義
├── dtos/           # リクエスト/レスポンスの型定義
└── infrastructure/ # DB接続・Claude APIクライアント
rdb/migrations/     # マイグレーションファイル（golang-migrate）
```

---

## 3. コミットメッセージのルール

- **1行・日本語・シンプルに**書く

```
# フォーマット
<動詞>: <変更内容の要約>
```

### 動詞の使い分け

| 動詞         | 用途              |
|------------|-----------------|
| `追加`       | 新機能・新ファイルの追加    |
| `修正`       | バグ修正            |
| `更新`       | 既存機能の改善・変更      |
| `削除`       | ファイル・コードの削除     |
| `リファクタリング` | 動作を変えない構造改善     |
| `ドキュメント`   | README・設計書などの更新 |
| `テスト`      | テストコードの追加・修正    |

**例:**

```
追加: 昆虫一覧取得APIの実装
```

---

## 4. コードスタイルルール（Go）

### 全般

- `go fmt` / `goimports` に準拠したフォーマットを維持する
- `golangci-lint` の警告が出ないコードを書く
- エラーは必ず処理する（`_` で無視しない）

### 命名規則

| 対象       | ルール                                     | 例                           |
|----------|-----------------------------------------|-----------------------------|
| パッケージ名   | 小文字・単数形                                 | `controller`, `service`     |
| 関数・メソッド  | UpperCamelCase（公開）/ lowerCamelCase（非公開） | `GetInsects`, `buildPrompt` |
| 変数       | lowerCamelCase                          | `insectList`, `aiComment`   |
| 定数       | UpperCamelCase または ALL_CAPS             | `MaxRetryCount`             |
| インターフェース | 先頭に `I` をつける                            | `IInsectRepository`         |
| ファイル名    | スネークケース                                 | `insect_service.go`         |

### コメントの書き方

```go
// GetInsects は昆虫の一覧をDBから全件取得して返す
func (r *insectRepository) GetInsects(ctx context.Context) ([]model.Insect, error) {
// SQL でinsectsテーブルを全件取得
...
}
```

- 公開関数・メソッドには必ずコメントを付ける
- コメントは関数名から始める（`// 関数名 は...` の形式）
- 処理の意図が分かりにくい箇所にはインラインコメントを付ける

### インターフェース設計

各層のインターフェースは別ファイルで定義する。

```go
// insect_repository_interface.go
type InsectRepository interface {
GetInsects(ctx context.Context) ([]model.Insect, error)
GetInsectByID(ctx context.Context, id int64) (*model.Insect, error)
}
```

### エラーハンドリング

```go
// エラーは呼び出し元でラップして返す
insects, err := r.insectRepo.GetInsects(ctx)
if err != nil {
return nil, fmt.Errorf("GetInsects: %w", err)
}
```

---

## 5. APIレスポンス形式

設計書（`docs/04_api.md`）のフォーマットに必ず従う。

```json
// 正常レスポンス（配列）
[{ ... }, { ... }]

// 正常レスポンス（オブジェクト）
{ ... }

// エラーレスポンス
{
  "error": {
    "code": 400,
    "message": "..."
  }
}
```

---

## 6. Claude API 利用ルール

- モデルは **Claude Haiku 4.5** を使用する
- APIエラー時は **3回リトライ** し、それでも失敗した場合はデフォルトコメントを返す
- プロンプトは `docs/04_api.md` の「AIプロンプト設計」に従って実装する
- **APIキーは環境変数で管理し、コードにハードコードしない**

---

## 7. セキュリティへの配慮

- **APIキー・DB接続情報などの機密情報は環境変数で管理**する
- 機密情報を含むファイルは必ず `.gitignore` に追加する
- SQLは必ずプレースホルダーを使い、SQLインジェクションを防ぐ
- ユーザーからの入力値はController層でバリデーションする
- エラーレスポンスに内部情報（スタックトレース等）を含めない

---

## 8. コードの品質・設計方針

### やること

- 依存関係の注入（DI）を徹底し、各層はインターフェースに依存させる
- テストしやすい設計を維持する（Repository層のモック差し替えを可能にする）
- 処理の責務を各層に正しく割り当てる

| 層          | やること            | やらないこと          |
|------------|-----------------|-----------------|
| Controller | バリデーション・レスポンス整形 | ビジネスロジック・DB操作   |
| Service    | ビジネスロジック・AI呼び出し | 直接のDB操作・HTTPの詳細 |
| Repository | DB操作のみ          | ビジネスロジック        |

### やらないこと（過剰実装の禁止）

- 現時点で不要な将来機能の先取り実装
- 1箇所でしか使わない関数・インターフェースの抽象化
- 既存の動作に影響しない不要なリファクタリング
- 求められていないエラーハンドリングの追加

---

## 9. DB・マイグレーション

- スキーマ変更は必ず `rdb/migrations/` にマイグレーションファイルを追加する
- ファイル名は `NNNNNN_<説明>.up.sql` / `NNNNNN_<説明>.down.sql` のセットで作成
- ローカルでの実行: `make migrate-up` / `make migrate-down`
- カラムコメントは日本語で記述する（`docs/03_database.md` 参照）

---

## 10. 参照すべき設計書

| ドキュメント     | パス                        | 参照タイミング            |
|------------|---------------------------|--------------------|
| 要件定義書      | `docs/01_requirements.md` | 機能の仕様確認            |
| アーキテクチャ設計書 | `docs/02_architecture.md` | 構成・技術スタックの確認       |
| データベース設計書  | `docs/03_database.md`     | テーブル・カラム定義の確認      |
| API設計書     | `docs/04_api.md`          | エンドポイント・レスポンス形式の確認 |
| OpenAPI仕様書 | `docs/openapi.yml`        | API仕様の詳細確認         |
