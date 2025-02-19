.SILENT:
.PHONY: create

create:
	migrate create -ext sql -dir migrations ${n}