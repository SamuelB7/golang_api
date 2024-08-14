build:
	@go build -o bin/GO_API cmd/main.go

run: build
	@./bin/GO_API

test:
	@go test -v ./...
