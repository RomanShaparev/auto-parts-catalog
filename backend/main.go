package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	catalogservice "auto-parts-catalog/backend/catalogservice/gen"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	{
		// Set up a connection to the server.
		conn, err := grpc.NewClient(os.Getenv("CATALOG_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := catalogservice.NewCountryServiceClient(conn)
		//c := catalogservice.NewWarehouseServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Country
		r, err := c.CreateCountry(ctx, &catalogservice.CreateCountryRequest{Name: "Россия"})
		//r, err := c.CreateCountry(ctx, &catalogservice.CreateCountryRequest{Name: "Казахстан"})
		//r, err := c.GetCountry(ctx, &catalogservice.GetCountryRequest{Id: 2})
		//r, err := c.ListCountries(ctx, &emptypb.Empty{})
		//r, err := c.DeleteCountry(ctx, &catalogservice.DeleteCountryRequest{Id: 1})

		// Warehouse
		//r, err := c.CreateWarehouse(ctx, &catalogservice.CreateWarehouseRequest{CountryId: 3, CityName: "Москва"})
		//r, err := c.ListWarehouses(ctx, &catalogservice.ListWarehousesRequest{CountryId: 1})

		if err != nil {
			log.Printf("ERROR: %v", err)
		} else {
			log.Printf("SUCCESS: %s", r)
		}

	}

	return
	{
		r := chi.NewRouter()
		r.Use(middleware.Logger)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world!!"))
		})

		http.ListenAndServe(":8080", r)
	}

}
