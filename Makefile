build:
	@go build -o bin/finance-backend ./cmd/server

run: build 
	@./bin/finance-backend

test:
	@go test ./...
