DB_USER=postgres
DB_PWD=kldsfn7jen
DB_NAME=furniprice
DB_HOST=127.0.0.1
DB_PORT=5555

build:
	go build -o server cmd/server/main.go

run-dev:
	DB_USER=$(DB_USER) \
	DB_PWD=$(DB_PWD) \
	DB_NAME=$(DB_NAME) \
	DB_HOST=$(DB_HOST) \
	DB_PORT=$(DB_PORT) \
	go run cmd/server/main.go

docker-db-build:
	docker build --build-arg DB_PWD=$(DB_PWD) --build-arg DB_NAME=$(DB_NAME) -t bsrv-test-furniprice-postgres-image ./build/db

docker-db-run:
	docker run -d --name bsrv-test-furniprice-postgres-name -p $(DB_PORT):5432 bsrv-test-furniprice-postgres-image
