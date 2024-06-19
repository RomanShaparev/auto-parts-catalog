.PHONY: migrate sqlc_generate grpc_generate

migrate:
	migrate -source file://catalog-service/postgres/migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

sqlc_generate:
	cd catalog-service; \
	sqlc generate

grpc_generate:
	protoc --go_out=./catalog-service/catalogservice/gen --go_opt=paths=source_relative \
		   --go-grpc_out=./catalog-service/catalogservice/gen --go-grpc_opt=paths=source_relative \
		   *.proto; \
	protoc --go_out=./backend/catalogservice/gen --go_opt=paths=source_relative \
		   --go-grpc_out=./backend/catalogservice/gen --go-grpc_opt=paths=source_relative \
		   *.proto