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
│   ├── cmd
│   │   ├── main.go
│   │   ├── wire.go
│   │   └── wire_gen.go
│   ├── config
│   │   ├── config.yaml
│   │   ├── database.go
│   │   ├── definition.go
│   │   └── router.go
│   ├── db
│   │   └── conf.d
│   ├── handlers
│   │   └── default.go
│   ├── log
│   │   └── mysql
│   ├── middlewares
│   ├── models
│   ├── repositories
│   │   └── default.go
│   ├── responses
│   │   ├── error.go
│   │   └── success.go
│   ├── services
│   │   └── default.go
│   └── test
│       ├── handler
│       │   └── default_test.go
│       └── helper.go
├── docker
│   └── mysql
│       ├── conf.d
│       │   └── mysql.conf
│       └── log
├── docker-compose.yml
├── go.mod
├── scripts
│   └── new_migration.sh
└── swagger
    ├── Dockerfile
    └── swagger.yml
```

# app_name

## Getting Started（開始方法）

migration ファイル作成のための brew

```shell
$ brew install golang-migrate
```

## Mysql, Swagger の立ち上げ

```
make up
```

swagger url: `http://localhost:3000/`

## Go サーバーの立ち上げ(ローカルの環境で)

```
make local
```

※docker 上で立ち上げたい場合は`docker-compose.yaml`のコメントアウトしてる箇所のコメントを外すことで docker でサーバーを立てることができる

## Tips

基本的な Tips は全て`Makefile`に記述してあるためそちらを参照すること

## その他

`default`というパッケージやファイルがあるが、それらは全て例として作成したものなので削除して使うこと。
