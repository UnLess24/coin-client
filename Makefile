.SILENT:
.PHONY: create up down lint swag

up:
	docker network create --driver bridge coinclientserver &> /dev/null || true
	docker-compose up -d

down:
	docker-compose down --rmi local
	docker network rm coinclientserver -f &> /dev/null || true

create:
	migrate create -ext sql -dir migrations ${n}

lint:
	golangci-lint run ./...

swag:
	swag init -g ./cmd/client/main.go -o cmd/docs
