.PHONY: migrate sqlc_generate grpc_generate

migrate:
	migrate -source file://catalog-service/storage/postgres/migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

sqlc_generate:
	cd catalog-service/storage; \
	sqlc generate

grpc_generate:
	cd catalog-service-proto; \
	protoc --go_out=./../catalog-service/api/grpc/gen --go_opt=paths=source_relative \
		   --go-grpc_out=./../catalog-service/api/grpc/gen --go-grpc_opt=paths=source_relative \
		   *.proto; \
	protoc --go_out=./../backend/catalogservice/grpc/gen --go_opt=paths=source_relative \
		   --go-grpc_out=./../backend/catalogservice/grpc/gen --go-grpc_opt=paths=source_relative \
		   *.proto