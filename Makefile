APP_NAME=hcm-be

.PHONY: tidy run build test lint fmt docker-build docker-up docker-down gen

tidy:
	@go mod tidy

run:
	@go run ./cmd/server

build:
	@go build -o bin/$(APP_NAME) ./cmd/server

test:
	@go test ./... -cover

fmt:
	@gofmt -s -w .
	@go vet ./...

lint:
	@golangci-lint run --timeout=5m

docker-build:
	@docker build -t $(APP_NAME):latest .

docker-up:
	@docker compose up -d

docker-down:
	@docker compose down

gen:
	@echo "place code generators here (e.g. mockgen, oapi-codegen)"
