# api/booking/Makefile
.PHONY: build run test docker migrate

SERVICE_NAME=booking
IMAGE_TAG=$(SERVICE_NAME):latest

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test -race -v ./...

docker:
	docker build -t $(IMAGE_TAG) .

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/*.proto
