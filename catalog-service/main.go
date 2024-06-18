package main

import (
	"auto-parts-catalog/catalog-service/catalogservice"
	"auto-parts-catalog/catalog-service/postgres/queries"
	"context"
	"flag"
	"log"
	"net"
	"os"
	"reflect"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_CONN_STRING"))
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := queries.New(conn)

	// list all authors
	authors, err := queries.ListCountries(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateCountry(ctx, "Томск")
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetCountry(ctx, insertedAuthor.CountryID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

// var (
// 	port = flag.Int("port", 50051, "The server port")
// )

// server is used to implement helloworld.GreeterServer.
type server struct {
	catalogservice.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *catalogservice.HelloRequest) (*catalogservice.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &catalogservice.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {

	flag.Parse()

	lis, err := net.Listen("tcp", ":"+os.Getenv("CATALOG_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	catalogservice.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
