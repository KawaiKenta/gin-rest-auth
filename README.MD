# Gin RestAPI auth

golang + gin で作成した REST API サーバー

## 技術構成

- go
- gin
- gorm
- mysql

## API

### 認証系統

- RESTful なバックエンドサーバーなのでログアウト機能は持ち得ないことに注意
  | | メソッド | URI | JSONdata |
  | :--------------- | :------- | :----------- | :---------------------------------------------------------------------- |
  | ユーザー新規登録 | POST | /auth/singup | {"email": "newuser5@gmail.com","name": "newbee","password": "password"} |
  | ユーザーログイン | POST | /auth/singin | {"email": "newuser5@gmail.com","password": "password"} |

### ユーザー CRUD

- Header に　 Authorization : Bearer {jwt_token}　を持つ時のみアクセス可能

  |                            | メソッド | URI        | JSONdata                                     |
  | :------------------------- | :------- | :--------- | :------------------------------------------- |
  | ユーザー情報の取得         | GET      | /v1/user   |                                              |
  | ユーザーの退会             | DELETE   | /v1/delete |                                              |
  | ユーザー情報のアップデート | POST     | /v1/update | {"password": "newpassword","name": "kkerew"} |

## 環境構築

### 1. データベース構築

- mysql のセットアップ
  cofig.ini と一致するように mysql データベースを作成すること

- マイグレーションを実行
- user テーブルを作成

```
go run tools/migrate/migrate.go
```

- シーディングを実行
- user データを作成

```
go run tools/seeds/seed.go
```

### 2. サーバー起動

- サーバーを起動する

```
go run main.go
```

### 3. データベースのテーブルドロップ

- user テーブルをドロップする

```
go run tools/clear/clear.go
```

## 開発ロードマップ

- [ ] カスタムバリデーションの実装
- [ ] ファイル構成の変更 (pakage にまとめる?)
