package main

import (
	"context"
	"fmt"
	"log"

	"hello-entgo/ent"
	"hello-entgo/ent/book"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/thanhpk/randstr"
)

func CreateBook(ctx context.Context, client *ent.Client) (*ent.Book, error) {
	u, err := client.Book.
		Create().
		SetNAME("The book").
		SetPRICE(39800).
		SetAUTHOR("Sir Author").
		SetISBN(randstr.String(8)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating book: %w", err)
	}
	log.Println("book was created: ", u)
	return u, nil
}

func QueryBook(ctx context.Context, client *ent.Client) (*ent.Book, error) {
	// `Only` fails if no book found, or more than 1 book returned.
	u, err := client.Book.Query().
		Where(book.NAMEEQ("The book")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying book: %w", err)
	}
	log.Println("book returned: ", u)

	return u, nil
}

func QueryBooks(ctx context.Context, client *ent.Client) ([]*ent.Book, error) {
	// `Only` fails if no book found, or more than 1 book returned.
	u, err := client.Book.Query().Where(book.NAMEEQ("The book")).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying book: %w", err)
	}
	log.Println("book returned: ", u)

	return u, nil
}

func main() {
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	// client, err := ent.Open("sqlite3", "file:bookshelf.db?_fk=1")
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "root", "pgsql", "bookshelf")
	// client, err := ent.Open("postgres", dsn)
	client, err := ent.Open("mysql", "root:@tcp(localhost:13306)/bookshelf")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Run the auto migration tool.
	err = client.Schema.Create(ctx)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	_, err = CreateBook(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	// _, err = QueryBook(ctx, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	_, err = QueryBooks(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
}
