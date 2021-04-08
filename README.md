# 使い方(setup)

```shell
make new
```

とすると
gin + gorm + mysql のディレクトリをざっくり作るよ

作成ソースツリー

```
.
├── Dockerfile
├── Makefile
├── README.md
├── app
│   ├── cmd
│   │   ├── main.go
│   │   ├── wire.go
│   │   └── wire_gen.go
│   ├── config
│   │   ├── config.yaml
│   │   ├── database.go
│   │   ├── definition.go
│   │   └── router.go
│   ├── db
│   │   └── conf.d
│   ├── handlers
│   │   ├── default.go
│   │   └── default_test.go
│   ├── log
│   │   └── mysql
│   ├── middlewares
│   ├── models
│   ├── repositories
│   │   └── default.go
│   ├── responses
│   │   ├── error.go
│   │   └── success.go
│   └── services
│       └── default.go
├── cover.html
├── cover.out
├── docker
│   └── mysql
│       ├── conf.d
│       │   └── mysql.conf
│       └── log
├── docker-compose.yml
├── go.mod
├── go.sum
├── scripts
│   └── new_migration.sh
└── swagger
    ├── Dockerfile
    └── swagger.yml
```

# app_name

## Getting Started（開始方法）

基本方針としては mysql などは全て docker-compose で立ち上げ、 メインサーバはそれぞれ docker run するようにする

```shell
#dbとかの実行
$ make up
```

```shell
# apiサーバの実行
$ make local
or
$ cd api && go run main.go
```

migration ファイル作成のための brew

```shell
$ brew install golang-migrate
```

### Prerequisites（前提条件）

使用言語やフレームワーク
必要なミドルウェアとそのインストール方法

### Installing

ローカル環境で動かすための手順を記述していく
最終的にどのような画面になれば適切に動かすことができているかを記述する

## Running the tests(テスト方法を記述)

テストを動かすための手順を記述する

## Deployment（デプロイ）

デプロイをする方法やどのような構成担っているかを説明する
