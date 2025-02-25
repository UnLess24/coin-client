# List of commands to start project

- print **make up** - to build and start project
- print **make down** - to finish project and remove all tmp containers and images
- print **make lint** - to easy linting project by golangci-lint
- print **make create n=$nameOfMigration** - to create migration

## Migrations commands

- build and run cmd/migrate/main.go to migrate migrations
- ./migrate up 1/migrate down 1 - to process up/down migration, number after command is optional
