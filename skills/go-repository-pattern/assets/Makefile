.PHONY: sqlc migrate-up migrate-down

sqlc:
	cd internal/{domain}/infrastructure/repository && sqlc generate

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1
