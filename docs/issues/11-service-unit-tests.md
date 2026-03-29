## 背景 / 目的

#5〜#10 の各 Issue で TDD によりテストを書きながら実装を進めた結果、
Service 層のテストカバレッジを計測し、抜け漏れがあれば補完する。
最終的に Service 層の品質を数値で担保する。

- 依存: #8, #10
- ラベル: `backend`, `test`

---

## スコープ / 作業項目

- `make test` でテスト全体を実行し、カバレッジレポートを出力
- カバレッジが 70% 未満の箇所に追加テストを書く:
  - `insect_service_test.go` の未カバーケース
  - `question_service_test.go` の未カバーケース
  - `diagnosis_service_test.go` の未カバーケース（スコアの境界値 0 / 2 など）
- `Makefile` に `make coverage`（HTML カバレッジレポート生成）を追加

**参照設計書**:
- `docs/01_requirements.md`（診断ロジック詳細・エラー処理仕様）

---

## ゴール / 完了条件（Acceptance Criteria）

- [ ] `make coverage` でカバレッジレポートが HTML で出力される
- [ ] Service 層全体のカバレッジが 70% 以上
- [ ] `DiagnosisService` のスコア境界値（0/0/0 と 2/2/2）がテストされている
- [ ] Claude API の1回目失敗 → リトライ成功のケースがテストされている
- [ ] `make test` がエラーなく通る

---

## テスト観点

- カバレッジ計測:
  ```bash
  # カバレッジ付きでテスト実行
  go test ./internal/service/... -cover

  # HTML レポート生成
  go test ./internal/service/... -coverprofile=coverage.out
  go tool cover -html=coverage.out -o coverage.html
  ```
- 検証方法: `coverage.html` をブラウザで開き、赤（未カバー）の行を確認して追加テストを書く

---

## 要確認事項

- `coverage.html` は `.gitignore` に追加する
- CI（GitHub Actions）でカバレッジを自動チェックするかは Phase 4 のデプロイ整備とあわせて検討
