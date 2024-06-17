.PHONY: migrate sqlc_generate

migrate:
	migrate -source file://catalog-service/migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

sqlc_generate:
	cd catalog-service; \
	sqlc generate