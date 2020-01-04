# clean-arch-go-ja

ロバート マーチンさんが提唱したクリーンアーキテクチャでAPIを実装したものです。

# ルール
* 外部デバイス (mysql, echo, s3など) に依存しないこと(交換可能であることを意味する)
* ビジネスルールがテスト可能であること

# フォルダ構成
```
.
├── cmd  # mainパッケージが入ってます。
│   ├── apiserver
│   │   └── main.go
│   └── db-migrate
│       └── main.go
├── internal
│   ├── domain
│   │   └── model  # ここに、モデルを書いていきます。 他の何にも依存してはいけません。
│   │       └── user.go
│   ├── infrastructure  # 外部デバイスを扱うためのパッケージです。基本的には、外部デバイスはここ以外では使うことは許されません。
│   │   ├── conf  # 設定を読み込むパッケージ
│   │   │   ├── conf.go
│   │   │   └── yaml
│   │   │       └── viper.go
│   │   ├── datastore  # データを保存する系のパッケージ
│   │   │   ├── db.go
│   │   │   ├── gorm
│   │   │   │   └── gorm.go
│   │   │   └── storage.go
│   │   └── web
│   │       ├── doc.go
│   │       ├── echo
│   │       │   ├── context.go
│   │       │   ├── group.go
│   │       │   ├── middleware.go
│   │       │   └── server.go
│   │       ├── echov4  # 交換可能であることを示すために、echov4も実装してます。
│   │       │   ├── context.go
│   │       │   ├── group.go
│   │       │   ├── middleware.go
│   │       │   └── server.go
│   │       └── server.go
│   ├── interface  # infra の形式から usecase で使いやすい形式に変換するためのパッケージ
│   │   ├── controller
│   │   │   ├── app_controller.go
│   │   │   ├── doc.go
│   │   │   └── user_controller.go
│   │   ├── repository
│   │   │   └── user_repository.go
│   │   └── web
│   │       ├── appcontext  # Contextに便利な関数を追加しとくことで、入力データの管理が楽になります。
│   │       │   ├── binder.go
│   │       │   ├── context.go
│   │       │   ├── doc.go
│   │       │   └── parameter.go
│   │       └── route
│   │           ├── doc.go
│   │           └── route.go
│   └── usecase  #  ビジネスルールをかく
│       └── user_usecase.go
```
