package main

import (
	"log"
	"os"

	"auto-parts-catalog/backend/api"
	"auto-parts-catalog/backend/catalogservice/grpcimpl"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Set up a connection to the catatalog service
	conn, err := grpc.NewClient(os.Getenv("CATALOG_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	catalogService := grpcimpl.NewCatalogService(conn)

	// Set up server
	server := api.NewServer(catalogService)
	err = server.Start(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
