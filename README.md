# 使い方

```shell
./new_gin.sh プロジェクト名 ディレクトリ名
```

とすると
gin + gorm + mysql のディレクトリをざっくり作るよ

作成ソースツリー

```
│
├── Makefile
├── README.md
├── api
│   ├── Dockerfile
│   ├── config
│   │   ├── database.go
│   │   └── util.go
│   ├── db
│   │   ├── conf.d
│   │   │   └── mysql.conf
│   │   └── migrations
│   ├── go.mod
│   ├── go.sum
│   ├── handler
│   ├── lib
│   ├── log
│   │   └── mysql
│   ├── main.go
│   ├── models
│   │   ├── repositories
│   │   └── services
│   └── responses
│       └── error.go
├── docker-compose.yml
├── new_gin.sh
├── swagger
│   ├── Dockerfile
│   └── swagger.yml
└── terraform
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
