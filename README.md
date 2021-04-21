# Practice & Taste

```
ent/ent - entgo
```


## Init

```sh
$ go mod init hello-entgo
$ go get entgo.io/ent/cmd/ent
$ go run entgo.io/ent/cmd/ent init Book
```


## Edit `Fields()` in `ent/schema/book.go`

```go
// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("NAME").Default("unknown"),
		field.Int("PRICE").Positive(),
		field.String("AUTHOR").Default("unknown"),
		field.String("ISBN").Unique().Default("unknown"),
	}
}
```


## Generate

```sh
$ go generate ./...
```

## Run - Database have to be created previously on `MySQL`, `Postgres`
```sh
$ go run .
```
