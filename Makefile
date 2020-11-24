new:
	./new_gin.sh

up:
	docker-compose up --build -d

down:
	docker-compose down

local:
	# docker build -t app_name-api ./api
	# docker run -e MYSQL_HOST=host.docker.internal --rm --name app_name-api_1 -p 8080:8080 app_name-api
	go run ./api/main.go ./api/wire_gen.go

# テストカバレッジをhtmlで表示する
cover:
	go test -cover ./... -coverprofile=cover.out.tmp
	# 自動生成コードをカバレッジ対象から外し、カバレッジファイルを作成
	cat cover.out.tmp | grep -v "wire_gen.go" > cover.out
	rm cover.out.tmp
	go tool cover -html=cover.out -o cover.html
	open cover.html

lint:
	golangci-lint run -E golint

ssh-api:
	docker exec -it app_name-api_1 /bin/bash

ssh-mysql:
	docker exec -it app_name-mysql /bin/bash

# terraform:
# 	terraform init
# 	terraform plan
	# applyは手動でやること
	# terraform apply

# new:
# 	./scripts/new_migration.sh

migrate-up:
	migrate -source file://api/db/migrations/ -database 'mysql://root:password@tcp(127.0.0.1:3306)/app_name' up

migrate-down:
	migrate -source file://api/db/migrations/ -database 'mysql://root:password@tcp(127.0.0.1:3306)/app_name' down

migrate-ver:
	migrate -source file://api/db/migrations/ -database 'mysql://root:password@tcp(127.0.0.1:3306)/app_name' version
