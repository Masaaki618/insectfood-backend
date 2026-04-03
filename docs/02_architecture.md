# アーキテクチャ設計書｜昆虫食初心者ガイド

上記の要件をベースに、個人開発向けのシンプルなアーキテクチャを設計する。

## 制約

- 個人開発（一人）
- 低コスト重視
- 運用が簡単

---

## 技術スタック

- **フロントエンド**: Flutter + Dart
- **バックエンド**: Go + Gin（Cloud Run）
- **データベース**: MySQL互換（TiDB Cloud Starter）
- **AI API**: Claude Haiku 4.5（昆虫レコメンド・コメント生成）
- **認証**: なし（MVP段階）
- **インフラ**: Cloud Run + TiDB Cloud

---

## システム構成図

```mermaid
graph TD
    A[Flutter App\niOS / Android] -->|HTTP Request| B[Router\nroute.go]
    B --> C[Controller]
    C --> D[Service]
    D --> E[Repository]
    E -->|Query| F[TiDB Cloud\nMySQL]
    F -->|Response| E
    E --> D
    D -->|スコア + 昆虫リスト| G[Claude Haiku 4.5\nAI API]
    G -->|レコメンド + コメント| D
    D --> C
    C -->|Response JSON| A
```

---

## アプリケーション構成

バックエンドは三層アーキテクチャで実装する。

| 層          | 役割               | 具体例                                  |
|------------|------------------|--------------------------------------|
| Router     | ルーティング（route.go） | エンドポイントの定義・Controllerへの振り分け          |
| Controller | リクエスト/レスポンスの処理   | リクエストのバリデーション・Serviceの呼び出し・レスポンスの返却  |
| Service    | ビジネスロジック         | カテゴリ別スコアの集計・Claude APIの呼び出し・レスポンスの整形 |
| Repository | DBアクセス・データの取得・保存 | 昆虫一覧の取得・質問のランダム取得                    |

各層はインターフェースを介してやり取りをする。GoのDuck Typingを活かし、各層がインターフェースに依存する設計とする。

これにより以下のメリットを得る。

- DBの変更が発生した場合、変更箇所をRepository層内に抑え込むことができ、ControllerやServiceへの影響を最小限に抑える
- 本番DB実装・インメモリ実装・モック実装を切り替え可能にすることでテストの容易性を確保する

---

## Flutter構成

| 項目         | ライブラリ     | 概要                 |
|------------|-----------|--------------------|
| 状態管理       | Riverpod  | 現在主流・テストしやすい・DI対応  |
| HTTPクライアント | dio       | API通信・エラーハンドリングが充実 |
| レーダーチャート   | fl_chart  | 昆虫詳細画面のレーダーチャート表示  |
| ルーティング     | go_router | 画面遷移の管理            |

**フォルダ構成**

```
lib/
├── main.dart
├── app.dart                        # アプリ全体の設定・ルーティング
│
├── features/                       # 機能ごとに分ける
│   ├── insects/                    # 昆虫関連
│   │   ├── data/
│   │   │   └── insect_repository.dart
│   │   ├── domain/
│   │   │   └── insect.dart         # モデル
│   │   └── presentation/
│   │       ├── insects_page.dart
│   │       ├── insect_detail_page.dart
│   │       └── insects_provider.dart
│   │
│   ├── diagnosis/                  # 診断関連
│   │   ├── data/
│   │   │   └── diagnosis_repository.dart
│   │   ├── domain/
│   │   │   └── diagnosis.dart
│   │   └── presentation/
│   │       ├── diagnosis_page.dart
│   │       ├── diagnosis_result_page.dart
│   │       └── diagnosis_provider.dart
│   │
│   └── questions/                  # 質問関連
│       ├── data/
│       │   └── question_repository.dart
│       ├── domain/
│       │   └── question.dart
│       └── presentation/
│           ├── question_page.dart
│           └── question_provider.dart
│
└── shared/                         # 共通部品
    ├── widgets/                    # 共通ウィジェット
    ├── theme/                      # テーマ・カラー
    └── api/                        # API通信の共通処理
        └── api_client.dart
```

---

## バックエンド構成（Go/Gin）

**フォルダ構成**

```
.
├── main.go
├── Dockerfile
├── docker-compose.yml
├── Makefile
│
├── cmd/
│   └── server/
│       └── main.go             # サーバー起動
│
├── internal/
│   ├── routers/
│   │   └── router.go           # ルーティング定義
│   │
│   ├── controllers/            # リクエスト/レスポンス処理
│   │   ├── insect_controller.go
│   │   ├── question_controller.go
│   │   └── diagnosis_controller.go
│   │
│   ├── services/               # ビジネスロジック
│   │   ├── insect_service.go
│   │   ├── insect_service_interface.go
│   │   ├── question_service.go
│   │   ├── question_service_interface.go
│   │   ├── diagnosis_service.go
│   │   └── diagnosis_service_interface.go
│   │
│   ├── repositories/           # DBアクセス
│   │   ├── insect_repository.go
│   │   ├── insect_repository_interface.go
│   │   ├── question_repository.go
│   │   └── question_repository_interface.go
│   │
│   ├── models/                 # DBのモデル定義
│   │   ├── insect.go
│   │   ├── radar_chart.go
│   │   └── question.go
│   │
│   ├── dtos/                   # リクエスト/レスポンスの型定義
│   │   ├── insect_dto.go
│   │   ├── question_dto.go
│   │   └── diagnosis_dto.go
│   │
│   └── infrastructure/         # 外部サービス連携
│       ├── database/
│       │   └── db.go           # DB接続
│       └── ai/
│           └── claude_client.go # Claude API クライアント
│
└── rdb/
    └── migrations/             # マイグレーションファイル
        ├── 000001_create_insects_table.up.sql
        ├── 000001_create_insects_table.down.sql
        ├── 000002_create_radar_charts_table.up.sql
        ├── 000002_create_radar_charts_table.down.sql
        ├── 000003_create_questions_table.up.sql
        └── 000003_create_questions_table.down.sql
```

| 項目       | ツール                         | 概要                     |
|----------|-----------------------------|------------------------|
| コンテナ管理   | Docker + docker-compose     | アプリ・DBをコンテナで管理         |
| ホットリロード  | air                         | コード変更を検知して自動再起動        |
| ORM        | GORM                        | DBアクセスをGoの構造体で管理        |
| マイグレーション | golang-migrate（Makefileで管理） | up/downセットで管理・ロールバック可能 |
| テスト実行    | ginkgo（Makefileで管理）         | 各層のユニットテスト実行           |

**マイグレーションの運用イメージ**

```
rdb/migrations/
├── 000001_create_insects_table.up.sql
├── 000001_create_insects_table.down.sql
├── 000002_create_questions_table.up.sql
└── 000002_create_questions_table.down.sql

# 実行コマンド
make migrate-up    # 前進
make migrate-down  # 巻き戻し
```

---

## 選定理由

- **Flutter**: 会社で使用しているBubbleのリプレイス候補として、先行して個人開発で技術検証するため
- **Go / Gin**: 会社のBubbleバックエンドがGin構築のため、既存の技術スタックに合わせて選定
- **MySQL互換（TiDB Cloud）**: 低コスト重視の制約に基づき、MySQL互換で無料枠が5GBと大きいTiDB Cloud Starterを選定
- **Cloud Run**: サーバーレスのためリクエストがない時のコストがゼロ。無料枠が月200万リクエストと大きく、100人規模の個人開発では十分カバーできるため選定

---

## 初期コスト（月額）

- Cloud Run: $0（無料枠内）
- TiDB Cloud Starter: $0（無料枠内・5GB）
- Claude Haiku 4.5 API: 約$0.1（100人全員が診断した場合の上限）
- **合計**: 約$0.1/月
