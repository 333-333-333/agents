.PHONY: build run test test-integration test-coverage docker sqlc migrate-up migrate-down

SERVICE_NAME={name}
IMAGE_TAG=$(SERVICE_NAME):latest
LOCAL_PORT={assigned port from table}

build:
	go build -o bin/server ./cmd/server

run:
	APP_ENV=local \
	HTTP_PORT=$(LOCAL_PORT) \
	DB_PROVIDER=memory \
	go run ./cmd/server

run-pg:
	APP_ENV=development \
	HTTP_PORT=$(LOCAL_PORT) \
	DB_PROVIDER=postgres \
	DB_HOST=localhost \
	DB_PORT={pg host port from table} \
	DB_USER={service_name} \
	DB_PASSWORD={service_name} \
	DB_NAME=bastet_{service_name} \
	DB_SSL_MODE=disable \
	go run ./cmd/server

test:
	go test -race -v ./...

test-integration:
	go test -race -v -tags=integration ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

sqlc:
	cd internal/{domain}/infrastructure/repository && sqlc generate

docker:
	docker build -t $(IMAGE_TAG) .

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1
