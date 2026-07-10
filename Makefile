.PHONY: build run tidy

build:
	go build -o bin/app main.go
run:
	go run cmd/main.go
tidy:
	go mod tidy