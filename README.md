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

## docs

## アーキテクチャ

ddd アーキテクチャ

### er 図

<img width="459" alt="スクリーンショット 2020-06-19 16.04.26.png" src="https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/281558/2c60b4b0-224d-7341-949c-154802369c41.png">

### api モック

https://stoplight.io/p/docs/gh/ca21engineer/app_name/reference/main.v1.yaml/components/schemas/Round_Score?srn=gh/ca21engineer/app_name/reference/main.v1.yaml/components/schemas/Round_Score&group=master
