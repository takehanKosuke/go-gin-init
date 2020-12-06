# !/bin/bash

read -p "migrationファイル名を入力してください: " file_name

migrate create -ext sql -dir docker/mysql/migrations -seq $file_name
