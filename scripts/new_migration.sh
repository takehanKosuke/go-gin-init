# !/bin/bash

read -p "migrationファイル名を入力してください: " file_name

migrate create -ext sql -dir api/db/migrations -seq $file_name
