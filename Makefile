build:
	@go build -o bin/GO_API cmd/main.go

run: build
	@./bin/GO_API

test:
	@go test -v -cover ./...

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations -seq $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down