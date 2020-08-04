DB_USER=postgres
DB_PWD=kldsfn7jen
DB_NAME=furniprice
DB_LOCALHOST_HOST=127.0.0.1
DB_LOCALHOST_PORT=5555
DOCKER_DB_PORT=5432
DOCKER_DB_CONTAINER=bsrv-test-furniprice-postgres
DOCKER_APP_CONTAINER=bsrv-test-furniprice-app
DOCKER_NETWORK=bsrv-test-furniprice-net
APP_PORT=8080

build:
	go build -o server cmd/server/main.go

run-dev:
	DB_USER=$(DB_USER) \
	DB_PWD=$(DB_PWD) \
	DB_NAME=$(DB_NAME) \
	DB_HOST=$(DB_LOCALHOST_HOST) \
	DB_PORT=$(DB_LOCALHOST_PORT) \
	go run cmd/server/main.go

all: docker-network-up docker-db-build docker-db-run docker-app-build docker-app-run

docker-network-up:
	docker network inspect $(DOCKER_NETWORK) >/dev/null 2>&1 || \
    docker network create --driver bridge $(DOCKER_NETWORK)

docker-app-build:
	docker build --build-arg APP_PORT=$(APP_PORT) \
		-t $(DOCKER_APP_CONTAINER)-image \
		-f ./build/package/app/Dockerfile ./

docker-app-run:
	docker run -d --name $(DOCKER_APP_CONTAINER) \
		-e DB_USER=$(DB_USER) \
		-e DB_PWD=$(DB_PWD) \
		-e DB_NAME=$(DB_NAME) \
		-e DB_HOST=$(DOCKER_DB_CONTAINER) \
		-e DB_PORT=$(DOCKER_DB_PORT) \
		-p $(APP_PORT):$(APP_PORT) \
		--network $(DOCKER_NETWORK) \
		$(DOCKER_APP_CONTAINER)-image

docker-db-build:
	docker build --build-arg DB_PWD=$(DB_PWD) \
		--build-arg DB_NAME=$(DB_NAME) \
		-t $(DOCKER_DB_CONTAINER)-image ./build/package/db

docker-db-run:
	docker run -d --name $(DOCKER_DB_CONTAINER) \
		-p $(DB_LOCALHOST_PORT):$(DOCKER_DB_PORT) \
		--network $(DOCKER_NETWORK) \
		$(DOCKER_DB_CONTAINER)-image

clean:
	sudo docker stop $(DOCKER_APP_CONTAINER) || true
	sudo docker stop $(DOCKER_DB_CONTAINER) || true
	sudo docker rm $(DOCKER_APP_CONTAINER) || true
	sudo docker rm $(DOCKER_DB_CONTAINER) || true
	sudo docker network rm $(DOCKER_NETWORK) || true
