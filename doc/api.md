# Catalyst プロジェクト API仕様書

## 1\. 概要

本APIは、プロジェクト『Catalyst』のバックエンド機能を提供します。通信プロトコルには**GraphQL**を採用し、フロントエンドが必要なデータを効率的に、かつ過不足なく取得できる設計となっています。

本ドキュメントは、`schema/*.graphqls`ファイルに定義されたスキーマの仕様を解説するものです。

---

## 2\. エンドポイント

GraphQL APIのエンドポイントは、通常、単一のURLとなります。

-   **URL**: `/graphql` (または `/query`)
    

すべてのリクエスト（Query, Mutation）は、このエンドポイントに対してPOSTメソッドで送信されます。

---

## 3\. 型定義

### 3.1. カスタムスカラー

-   **`Time`**: 日時情報を扱うためのカスタムスカラー型です。ISO 8601形式の文字列として扱われます。
    
-   **`URI`**: URI (Uniform Resource Identifier) を表すためのカスタムスカラー型です。URL形式の文字列として扱われます。
    

### 3.2. オブジェクト型

APIで取得できるデータの基本単位です。

#### `Task`

タスク情報を表す中心的なオブジェクトです。

```
type Task {
    id: ID!
    title: String!
    description: String
    status: TaskStatus!
    importance: Importance!
    dueDate: Time!
    messageURL: URI
    channelId: ID!
    category: Category!
    assignees: [User!]!
    reminders: [Reminder!]!
    createdAt: Time!
    updatedAt: Time!
}
```

#### `Category`

タスクの分類を表すオブジェクトです。

```
type Category {
    id: ID!
    name: String!
}
```

#### `User`

担当者（traQユーザー）を表すオブジェクトです。


```
type User {
    traqId: ID! # traQ User UUID
    name: String!
    icon: String
}
```

#### `Reminder`

リマインダー設定を表すオブジェクトです。

```
type Reminder {
    id: ID!
    remindAt: Time!
}
```

### 3.3. Enum (列挙型)

特定の値しか取らないフィールドに使用される型です。

-   **`TaskStatus`**: タスクの進行状況を表します。
    
    -   `IN_PROGRESS`: 進行中
        
    -   `COMPLETED`: 完了
        
    -   `WAITING`: 返信待ち
        
-   **`Importance`**: タスクの重要度を表します。
    
    -   `HIGH`: 高
        
    -   `MEDIUM`: 中
        
    -   `LOW`: 低
        
-   **`TaskSortKey`**: タスクの並び替え基準を表します。
    
    -   `DUE_DATE`: 期限日
        
    -   `CREATED_AT`: 作成日
        
    -   `NEXT_REMINDER_AT`: 直近のリマインド日時
        
-   **`SortOrder`**: 並び順を表します。
    
    -   `ASC`: 昇順
        
    -   `DESC`: 降順
        

### 3.4. Input (入力型)

データの作成・更新（Mutation）や、複雑なクエリの引数をまとめて渡すための型です。

#### `CreateTaskInput`

新規タスク作成時に使用します。

```
input CreateTaskInput {
    title: String!
    description: String
    status: TaskStatus = IN_PROGRESS
    importance: Importance!
    dueDate: Time!
    messageURL: URI
    channelId: ID!
    categoryId: ID!
    assigneeIds: [ID!]!
    reminderDates: [Time!]
}
```

#### `UpdateTaskInput`

既存タスク更新時に使用します。

```
input UpdateTaskInput {
    id: ID!
    title: String
    description: String
    status: TaskStatus
    importance: Importance
    dueDate: Time
    messageURL: URI
    channelId: ID
    categoryId: ID
    assigneeIds: [ID!]
}
```

#### `CreateCategoryInput`

新規カテゴリ作成時に使用します。

```
input CreateCategoryInput {
    name: String!
}
```

#### `UpdateCategoryInput`

既存カテゴリ更新時に使用します。

```
input UpdateCategoryInput {
    id: ID!
    name: String!
}
```

#### `TaskFilterInput`

タスク一覧を絞り込む条件を指定します。

```
input TaskFilterInput {
    status: [TaskStatus!]
    importance: [Importance!]
    assigneeIds: [ID!]
    dueDateBefore: Time
    dueDateAfter: Time
}
```

#### `TaskSortInput`

タスク一覧を並び替える条件を指定します。

```
input TaskSortInput {
    key: TaskSortKey!
    order: SortOrder = DESC
}
```

---

## 4\. 操作一覧

### 4.1. Query (データ取得)

-   **`tasks(filter: TaskFilterInput, sortBy: TaskSortInput): [Task!]!`**
    
    -   **説明**: タスクの一覧を取得します。引数なしで実行した場合、デフォルトの条件でソートされたタスクが返されます。`filter`や`sortBy`引数を使用することで、「進行中のタスクのみを重要度で絞り込み、期限が近い順に並び替える」といった柔軟なデータ取得が可能です。
        
-   **`task(id: ID!): Task`**
    
    -   **説明**: 指定した単一のタスクに関する全ての詳細情報を取得します。タスクの詳細表示画面や編集画面の初期表示で利用することを想定しています。指定したIDのタスクが存在しない場合は `null` が返されます。
        
-   **`categories: [Category!]!`**
    
    -   **説明**: 登録されている全てのカテゴリを取得します。主に、タスク作成・編集フォームのカテゴリ選択用ドロップダウンリストを生成するために使用します。
        

### 4.2. Mutation (データ変更)

-   **`createTask(input: CreateTaskInput!): Task!`**
    
    -   **説明**: `CreateTaskInput`オブジェクトに含まれる情報（タイトル、期限、担当者など）を元に、新しいタスクをデータベースに作成します。成功すると、サーバーで採番された`id`を含む、作成されたばかりの`Task`オブジェクトが返されます。
        
-   **`updateTask(input: UpdateTaskInput!): Task!`**
    
    -   **説明**: `UpdateTaskInput`に含まれる`id`でタスクを特定し、指定されたフィールドの値で情報を更新します。`input`オブジェクトで省略されたフィールドは変更されません。成功すると、更新後の完全な`Task`オブジェクトが返されます。
        
-   **`deleteTask(id: ID!): ID!`**
    
    -   **説明**: 指定した`id`を持つタスクをデータベースから完全に削除します。操作が成功したことを確認するため、削除されたタスクの`id`が返されます。
        
-   **`createCategory(input: CreateCategoryInput!): Category!`**
    
    -   **説明**: タスクで使用するための新しいカテゴリを作成します。カテゴリ名はユニークである必要があります。成功すると、サーバーで採番された`id`を含む、作成された`Category`オブジェクトが返されます。
        
-   **`updateCategory(input: UpdateCategoryInput!): Category!`**
    
    -   **説明**: `id`で指定した既存のカテゴリの名前を変更します。成功すると、名前が変更された`Category`オブジェクトが返されます。