APP_NAME=go-uts

.PHONY: run build test migrate fresh seed

run:
	go run ./cmd/server

build:
	go build -o bin/$(APP_NAME) ./cmd/server

test:
	go test ./...

migrate:
	go run ./cmd/migrate -command up

fresh:
	go run ./cmd/migrate -command fresh

seed:
	go run ./cmd/seed
