APP_NAME=go-uts2
DOCKER_IMAGE=go-uts

.PHONY: run build test migrate fresh seed docker-build docker-run

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

docker-build:
	docker build -t $(DOCKER_IMAGE):latest .

docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE):latest
