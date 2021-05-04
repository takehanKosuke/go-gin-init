# !/bin/bash

# dockerクライアントの認証
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 666255764982.dkr.ecr.ap-northeast-1.amazonaws.com

# build
docker build -t go-gin-init .

# tagつけ（本当はtimestamp使ってやりたい）
docker tag go-gin-init:latest 666255764982.dkr.ecr.ap-northeast-1.amazonaws.com/go-gin-init:latest

# ecrリポジトリにプッシュ
docker push 666255764982.dkr.ecr.ap-northeast-1.amazonaws.com/go-gin-init:latest
