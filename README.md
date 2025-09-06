# Gin REST API

Go (Gin) + PostgreSQL を使用したユーザー管理REST APIです。

## 概要

このプロジェクトは、Gin WebフレームワークとGORM、PostgreSQLを使用して実装されたRESTful APIのデモアプリケーションです。ユーザーの作成、取得、更新、削除（CRUD操作）を提供します。

## API仕様

### エンドポイント

| メソッド | パス | 説明 |
|---------|------|------|
| GET | `/users` | 全ユーザー取得 |
| GET | `/users/:id` | 特定ユーザー取得 |
| POST | `/users` | ユーザー作成 |
| PUT | `/users/:id` | ユーザー更新（全項目） |
| PATCH | `/users/:id` | ユーザー更新（部分） |
| DELETE | `/users/:id` | ユーザー削除 |

## 環境構築

### 前提条件
- Docker & Docker Compose

### 1. アプリケーション起動

```bash
# リポジトリをクローン
git clone <repository-url>
cd <project-directory>

`.env` ファイルを作成・記述

# Docker Compose でアプリケーションを起動
docker-compose up --build
```

### 2. CRUD操作のテスト

1. **作成**
   ```bash
   curl -X POST http://localhost:3001/users \
     -H "Content-Type: application/json" \
     -d '{"name":"テストユーザー","email":"test@example.com","age":30}'
   ```

2. **取得**
   ```bash
   curl http://localhost:3001/users/1
   ```

3. **更新**
   ```bash
   curl -X PATCH http://localhost:3001/users/1 \
     -H "Content-Type: application/json" \
     -d '{"age":31}'
   ```

4. **削除**
   ```bash
   curl -X DELETE http://localhost:3001/users/1
   ```
