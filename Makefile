setup:
	@go mod tidy

run:
	@go build -o script-hub ./cmd/main.go
	@./script-hub

test:
	@go test ./...
