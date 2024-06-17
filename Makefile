.PHONY: migrate

migrate:
	migrate -source file://catalog-service/migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up