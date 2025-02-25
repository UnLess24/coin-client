.SILENT:
.PHONY: create up down lint

up:
	docker-compose up -d

down:
	docker-compose down --rmi local

create:
	migrate create -ext sql -dir migrations ${n}

lint:
	golangci-lint run ./...