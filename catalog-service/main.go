package main

import (
	"auto-parts-catalog/catalog-service/postgres/queries"
	"context"
	"log"
	"os"
	"reflect"

	"github.com/jackc/pgx/v5"
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

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
