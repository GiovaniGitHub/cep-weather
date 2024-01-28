CONTAINER_NAME := cep-weather
include .env

test:
	go mod tidy
	go test cmd/server/server_test.go

run:
	go mod tidy
	go run cmd/server/server.go

all:
	make test
	make run

build-docker:
	@echo "Construindo a imagem Docker..."
	@docker build -t ${CONTAINER_NAME} -f ./Dockerfile .

run-docker:
	@echo "Executando contêiner Docker..."
	@docker run -p ${WEB_SERVER_PORT}:${WEB_SERVER_PORT} ${CONTAINER_NAME}:latest
