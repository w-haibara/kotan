.PHONY: build
build: lint
	go build -o kotan main.go

.PHONY: lint
lint:
	go mod tidy
	go fmt ./...
	go vet ./...
	golangci-lint run
	gosec ./...
