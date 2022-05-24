# 使い方(setup)

```shell
make new
```

とすると
gin + gorm + mysql のディレクトリをざっくり作るよ

作成ソースツリー

```
.
├── app
│   ├── cmd
│   ├── config
│   ├── db
│   │   └── migrations
│   ├── handler
│   ├── middleware
│   │   └── mock_middleware
│   ├── repository
│   │   └── mock_repository
│   ├── response
│   ├── service
│   └── test
│       └── handler
├── docker
│   ├── mysql
│   │   ├── conf.d
│   │   └── log
│   └── swagger
├── proto
├── scripts
└── terraform
```

# app_name

## Getting Started（開始方法）

migration ファイル作成のための brew

```
brew install golang-migrate
```

mockgen の install

```
go install github.com/golang/mock/mockgen@v1.5.0
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

## swaggerの変更について

swaggerは `stoplight studio`を用いて記述していくことを前提としているためあえて1ファイルに全てのデータの記述を行なっている

```
brew install --cask stoplight-studio
```

## cloud run でのデプロイ

1. コンソール画面から新規プロジェクト`app_name`を作成
   https://console.cloud.google.com/projectselector2/home/dashboard

2. cloud run api を有効化

3. `./terraform/variable.tf` の project 部分を作成した project id に書き換える

4. `make cloud_run_apply` をすることで
   - cloud run
   - cloud sql
   - container registry
     の 3 つを作成する

## エラー解決

`make proto` とした時に

```
Please specify a program using absolute path or make sure the program is available in your PATH system variable」と言うエラーが出る場合
```

と言うようなエラーが吐かれた時は

```
export GOPATH=$HOME/go
PATH=$PATH:$GOPATH/bin
```

を.zshrc に追加することで解決する

## その他

`default`というパッケージやファイルがあるが、それらは全て例として作成したものなので削除して使うこと。
