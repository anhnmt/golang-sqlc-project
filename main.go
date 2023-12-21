package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/anhnmt/golang-sqlc-project/postgresql"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	dsn := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", "123456aA@"),
		Host:   fmt.Sprintf("%s:%d", "localhost", 5432),
		Path:   "sqlc",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	config, err := pgxpool.ParseConfig(dsn.String())
	if err != nil {
		return err
	}

	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(newCtx, config)
	if err != nil {
		return err
	}
	defer pool.Close()

	err = pool.Ping(newCtx)
	if err != nil {
		return err
	}

	queries := postgresql.New(pool)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// // create an author
	// insertedAuthor, err := queries.CreateAuthor(ctx, postgresql.CreateAuthorParams{
	//     Name: "Brian Kernighan",
	//     Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	// })
	// if err != nil {
	//     return err
	// }
	// log.Println(insertedAuthor)
	//
	// // get the author we just inserted
	// fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	// if err != nil {
	//     return err
	// }
	//
	// // prints true
	// log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}
