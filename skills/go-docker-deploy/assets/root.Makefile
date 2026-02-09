# Makefile (project root)
.PHONY: up down logs build-all test-all

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

build-all:
	docker-compose build

test-all:
	cd api/booking && go test -race ./...
	cd api/caregiver && go test -race ./...

migrate-all:
	cd api/booking && make migrate-up
	cd api/caregiver && make migrate-up
