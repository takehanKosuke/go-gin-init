up:
	docker-compose up --build -d

down:
	docker-compose down

local:
	# docker build -t app_name-api ./api
	# docker run -e MYSQL_HOST=host.docker.internal --rm --name app_name-api_1 -p 8080:8080 app_name-api
	go run ./api/main.go

ssh-api:
	docker exec -it app_name-api_1 /bin/bash

ssh-mysql:
	docker exec -it app_name-mysql /bin/bash
