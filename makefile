include .env.dev
export

dev:
	go run cmd/rest/main.go

lint:
	golangci-lint run

test:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

genmock:
	mockery --all
