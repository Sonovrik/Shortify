
NAME=auth
COMPOSE_FILE=docker-compose.yml
ENV=.env
PATH_TO_BUILD=./

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

docker:
	docker-compose -f $(PATH_TO_BUILD)$(COMPOSE_FILE) --env-file $(PATH_TO_BUILD)$(ENV) up --build -d

cleanDocker:
	docker-compose -f $(PATH_TO_BUILD)$(COMPOSE_FILE) --env-file $(PATH_TO_BUILD)$(ENV) down

fcleanDocker:
	docker system prune -f -a --volumes

.PHONY: build run test lint clean
