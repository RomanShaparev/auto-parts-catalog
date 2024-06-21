package main

import (
	grpcimpl "auto-parts-catalog/catalog-service/api/grpc"
	"auto-parts-catalog/catalog-service/api/grpc/gen"
	"auto-parts-catalog/catalog-service/storage/postgres"
	"auto-parts-catalog/catalog-service/storage/postgres/queries"
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
	storage := postgres.NewStorage(queries)

	// setup GRPC server
	lis, err := net.Listen("tcp", ":"+os.Getenv("CATALOG_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gen.RegisterCountryServiceServer(s, grpcimpl.NewCountryServiceServer(storage.CountryStorage))
	gen.RegisterWarehouseServiceServer(s, grpcimpl.NewWarehouseServiceServer(storage.WarehouseStorage))
	gen.RegisterCarModelServiceServer(s, grpcimpl.NewCarModelServiceServer(storage.CarModelStorage))
	gen.RegisterAutoPartServiceServer(s, grpcimpl.NewAutoPartServiceServer(storage.AutoPartStorage))
	gen.RegisterAutoPartComponentServiceServer(s, grpcimpl.NewAutoPartComponentServiceServer(storage.AutoPartComponentStorage))
	gen.RegisterWarehousePositionServiceServer(s, grpcimpl.NewWarehousePositionServiceServer(storage.WarehousePositionStorage))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
