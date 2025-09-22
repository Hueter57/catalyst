# Catalyst プロジェクト DB設計

## 概要

本プロジェクト『Catalyst』で使用するデータベースのスキーマ設計を以下に定義します。

## テーブル定義

### `task` テーブル

タスク本体の情報を格納します。

| カラム名 | データ型 | 制約 | 説明 |
| :--- | :--- | :--- | :--- |
| `id` | `BIGINT` | `PRIMARY KEY`, `AUTO_INCREMENT` | タスクの一意なID |
| `title` | `VARCHAR(255)` | `NOT NULL` | タスク名 |
| `description` | `TEXT` | `NULL` | タスクの詳細説明 |
| `status` | `Enum` | `NOT NULL`, `DEFAULT 'in_progress'` | ステータス (`in_progress`, `completed`, `waiting`) |
| `importance` | `Enum` | `NOT NULL` | 重要度 (`high`, `medium`, `low`) |
| `due_date` | `DATETIME` | `NOT NULL` | 期限日時 |
| `message_id` | `VARCHAR(36)` | `NULL` | 関連traQメッセージID |
| **`channel_id`** | **`VARCHAR(36)`** | **`NOT NULL`** | **通知先のtraQチャンネルUUID** |
| `created_at` | `DATETIME` | `NOT NULL`, `DEFAULT CURRENT_TIMESTAMP` | 作成日時 |
| `updated_at` | `DATETIME` | `NOT NULL`, `DEFAULT CURRENT_TIMESTAMP` | 更新日時 |
| `category_id` | `BIGINT` | `FOREIGN KEY` (category.id) | 外部キー (カテゴリ) |

### `category` テーブル

タスクに付与するカテゴリ情報を格納します。

| カラム名 | データ型 | 制約 | 説明 |
| :--- | :--- | :--- | :--- |
| `id` | `BIGINT` | `PRIMARY KEY`, `AUTO_INCREMENT` | カテゴリの一意なID |
| `name` | `VARCHAR(255)` | `NOT NULL`, `UNIQUE` | カテゴリ名 |

### `reminder` テーブル

タスクごとのリマインダー設定日時を格納します。

| カラム名 | データ型 | 制約 | 説明 |
| :--- | :--- | :--- | :--- |
| `id` | `BIGINT` | `PRIMARY KEY`, `AUTO_INCREMENT` | リマインダーの一意なID |
| `task_id` | `BIGINT` | `FOREIGN KEY` (task.id) | どのタスクのリマインダーか |
| `remind_at` | `DATETIME` | `NOT NULL` | リマインドを通知する日時 |

### `assignee` テーブル

タスクと担当者を紐付けるための中間テーブルです。これにより、タスクと担当者の「多対多」の関係を実現します。

| カラム名 | データ型 | 制約 | 説明 |
| :--- | :--- | :--- | :--- |
| `task_id` | `BIGINT` | `PRIMARY KEY`, `FOREIGN KEY` (task.id) | タスクのID |
| `user_id` | `VARCHAR(36)` | `PRIMARY KEY` | 担当者のtraQ UUID |
