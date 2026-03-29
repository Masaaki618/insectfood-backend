## 背景 / 目的

診断フロー画面（`/diagnosis`）が使用する `GET /api/v1/questions` エンドポイントを実装する。
questions テーブルからカテゴリ別に2問ずつランダムで抽出し、毎回異なる6問の組み合わせを返す。
**TDDで進める**。

- 依存: #4
- ラベル: `backend`

---

## スコープ / 作業項目

TDD の **Red → Green → Refactor** サイクルで進める。

### 🔴 Red（テストを先に書く）

- `model/question.go` — questions テーブルのモデル定義
- `repository/question_repository_interface.go` — `GetRandomQuestions()` インターフェース定義
- `service/question_service_interface.go` — インターフェース定義
- `repository/mock/mock_question_repository.go` — mockgen で自動生成
- `service/question_service_test.go` — 以下のテストケースを先に書く（この時点では失敗する）:
  - 正常系: Repository が6問返したとき、Service も6問返す
  - 正常系: カテゴリ別の件数が visual:2 / physical:2 / mental:2 になっている
  - 異常系: Repository がエラーを返したとき、Service もエラーを返す

### 🟢 Green（テストが通る実装を書く）

- `repository/question_repository.go` — カテゴリ別ランダム取得の実装（`ORDER BY RAND() LIMIT 2`）
- `service/question_service.go` — Repository を呼び出す実装
- `controller/question_controller.go` — レスポンス整形
- `dto/question_dto.go` — レスポンス用 DTO 定義
- `router/router.go` に `GET /api/v1/questions` のルートを追加

### 🔵 Refactor（コードを整理する）

- 重複があれば共通化（テストは引き続き通る状態を維持）

**参照設計書**:
- `docs/04_api.md`（質問取得APIのレスポンス形式）
- `docs/03_database.md`（questions テーブル定義・カテゴリ定義・初期質問12問）
- `docs/01_requirements.md`（診断ロジック: カテゴリ別2問ずつ計6問のルール）
- `docs/openapi.yml`（`/questions` スキーマ定義）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `make test` で `question_service_test.go` のテストがすべて通る（Green）
- [ ] `GET /api/v1/questions` が6問（visual: 2問 / physical: 2問 / mental: 2問）を返す
- [ ] 同じエンドポイントを複数回叩いたとき、異なる問題の組み合わせが返る（ランダム性の確認）
- [ ] レスポンスの各問題に `id`, `body`, `category` フィールドが含まれ `docs/04_api.md` の形式と一致している
- [ ] DBの各カテゴリに2問未満のデータしかない場合、適切なエラーを返す

---

## テスト観点

- ユニットテスト（ginkgo + gomock）:
  ```bash
  go test ./internal/service/... -v
  ```
  - 正常系: mock が6問（カテゴリ別2問ずつ）返すとき Service も6問返す
  - 異常系: mock がエラーを返すとき Service もエラーを返す
- 手動確認（curl）:
  ```bash
  # カテゴリ別件数の確認
  curl -s http://localhost:8080/api/v1/questions \
    | jq '[.data[].category] | group_by(.) | map({(.[0]): length}) | add'
  # {"mental":2,"physical":2,"visual":2}
  ```
