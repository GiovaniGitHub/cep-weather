test:
	go test cmd/server/server_test.go

run:
	go run cmd/server/server.go

all:
	make test
	make run