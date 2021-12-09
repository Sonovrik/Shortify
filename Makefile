.PHONY: build run test lint clean

NAME=auth

build:
	go build -v ./cmd/shortify.go

run:
	go run ./cmd/shortify.go

test:
	go test -v -timeout 30s ./...

lint:
	gofmt -s -w .
	golangci-lint run ./... --fix

clean:
	rm -rf $(NAME)

cleanSum:
	rm -rf go.sum
