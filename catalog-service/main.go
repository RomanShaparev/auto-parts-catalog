package main

import (
	"auto-parts-catalog/catalog-service/catalogservice"
	"auto-parts-catalog/catalog-service/catalogservice/gen"
	"auto-parts-catalog/catalog-service/postgres/queries"
	"auto-parts-catalog/catalog-service/storage"
	"context"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	// connect to DB
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_CONN_STRING"))
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer conn.Close(ctx)
	queries := queries.New(conn)

	// init Storage
	storage := storage.New(queries)

	// setup GRPC server
	lis, err := net.Listen("tcp", ":"+os.Getenv("CATALOG_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gen.RegisterCountryServiceServer(s, catalogservice.NewCountryServiceServer(storage.CountryStorage))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
