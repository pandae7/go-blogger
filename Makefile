.PHONY: server client test tidy

server:
	go run ./cmd/server

client:
	go run ./cmd/client

test:
	go test ./...

tidy:
	go mod tidy