.PHONY: test test-unit test-integration test-coverage

test: test-unit

test-unit:
	go test -race -v ./...

test-integration:
	go test -race -v -tags=integration ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html
