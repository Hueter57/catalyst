# Catalyst プロジェクト ディレクトリ構成案

[Catalist要件定義書]()

## 概要

本ドキュメントは、プロジェクト『Catalyst』の最終的なディレクトリ構成と、各ディレクトリ・ファイルが持つ役割を定義するものです。

## ディレクトリ構成と役割

```plaintext
catalyst/                                # プロジェクト全体のルートディレクトリ
├── backend/                             # Goで実装されたバックエンドアプリケーション
│   ├── cmd/                             # アプリケーションの起動コマンドを配置
│   │   └── server/                      # サーバーのメインパッケージ
│   │       └── main.go                  # ◆ アプリケーションのエントリーポイント（起動処理）
│   ├── internal/                        # 内部パッケージ（外部からの直接インポートを禁止）
│   │   ├── bot/                         # ◆ Bot関連のロジックを集約
│   │   │   ├── router/                  # Botのイベントを適切なハンドラに振り分けるルーター
│   │   │   ├── handler/                 # Botの各イベント（メンション等）を処理するハンドラ群
│   │   │   └── client/                  # traQ APIへの接続やメッセージ投稿を行うクライアント
│   │   ├── domain/                      # ◆ ビジネスロジックの核となるドメイン層
│   │   │   └── task.go                  # Task, Categoryなどのデータ構造を定義
│   │   ├── ent/                         # ◆ ORM `ent` が自動生成するコード
│   │   ├── graph/                       # ◆ `gqlgen` が自動生成するコード
│   │   │   └── model/                   # GraphQLのモデル（構造体）が生成される場所
│   │   ├── repository/                  # ◆ DBとの通信を担当するリポジトリ層
│   │   │   └── task.go                  # entを使い、DBの読み書きを実装
│   │   ├── resolver/                    # ◆ GraphQL APIの入り口となるリゾルバ層
│   │   │   └── resolver.go              # GraphQLのクエリに応じてServiceを呼び出す
│   │   └── service/                     # ◆ アプリケーションのビジネスロジックを担うサービス層
│   │       └── task.go                  # コアロジックを実装
│   ├── gqlgen.yml                       # `gqlgen` の設定ファイル
│   ├── go.mod                           # Goのモジュール定義ファイル
│   └── README.md                        # バックエンド固有の説明書
│
├── docs/                                # プロジェクトの設計書などを格納
│   └── architecture.md                  # アーキテクチャ図や設計思想など
│
├── frontend/                            # Vite + React (TypeScript) のフロントエンドアプリケーション
│   ├── public/                          # 静的ファイル（画像など）の配置場所
│   ├── src/                             # ソースコードのルート
│   │   ├── graphql/                     # ◆ GraphQL関連のコード
│   │   │   ├── client.ts                # GraphQLクライアントの初期化設定
│   │   │   ├── generated.ts             # スキーマから自動生成された型定義やHooks
│   │   │   └── queries.ts               # クエリやミューテーションの定義
│   │   ├── assets/                      # CSSから読み込む画像やフォントなど
│   │   ├── components/                  # ◆ 再利用可能なUIコンポー
│   │   │   ├── features/                # 特定の機能に紐づくコンポーネント
│   │   │   └── ui/                      # 汎用的なUIパーツ
│   │   ├── hooks/                       # ◆ カスタムReactフック
│   │   ├── pages/                       # ◆ 各ページに対応するトップレベルコンポーネント
│   │   ├── styles/                      # グローバルなCSSやテーマ設定
│   │   ├── types/                       # アプリケーション全体で使うTypeScriptの型定義
│   │   ├── utils/                       # 汎用的なヘルパー関数
│   │   ├── App.tsx                      # アプリケーションのルートコンポーネント
│   │   └── main.tsx                     # フロントエンドのエントリーポイント
│   ├── index.html                       # アプリケーションのベースとなるHTMLファイル
│   ├── package.json                     # パッケージ管理ファイル
│   ├── tsconfig.json                    # TypeScriptの設定ファイル
│   ├── vite.config.ts                   # Viteの設定ファイル
│   └── README.md                        # フロントエンド固有の説明書
│
├── schema/                              # ◆ APIの設計図となるGraphQLスキーマ
│   └── schema.graphqls                  # このファイルを元にバックエンドとフロントエンドのコードが生成される
│
└── README.md                            # プロジェクト全体の概要やセットアップ方法を記述
```
