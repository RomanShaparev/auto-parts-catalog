package main

import (
	"context"
	"flag"
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
		flag.Parse()
		// Set up a connection to the server.
		conn, err := grpc.NewClient(os.Getenv("CATALOG_SERVICE_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := catalogservice.NewCountryServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		r, err := c.CreateCountry(ctx, &catalogservice.CreateCountryRequest{Name: "Tomsk"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetName())

	}

	//return
	{
		r := chi.NewRouter()
		r.Use(middleware.Logger)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world!!"))
		})

		http.ListenAndServe(":8080", r)
	}

}
