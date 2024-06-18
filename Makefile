.PHONY: migrate sqlc_generate grpc_generate

migrate:
	migrate -source file://catalog-service/migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

sqlc_generate:
	cd catalog-service; \
	sqlc generate

grpc_generate:
	protoc --go_out=./catalog-service/catalogservice --go_opt=paths=source_relative \
		   --go-grpc_out=./catalog-service/catalogservice --go-grpc_opt=paths=source_relative \
		   *.proto; \
	protoc --go_out=./backend/catalogservice --go_opt=paths=source_relative \
		   --go-grpc_out=./backend/catalogservice --go-grpc_opt=paths=source_relative \
		   *.proto