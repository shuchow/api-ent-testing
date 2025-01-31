package main

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuchow/api-ent-testing/ent"
	"github.com/shuchow/api-ent-testing/handlers"

	"log"
	"net/http"
)

func main() {

	entClient, err := ent.Open(dialect.MySQL, "dbUser:dbPassword@tcp(db:3306)/dbName?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}

	defer entClient.Close()

	if err := entClient.Schema.Create(context.Background()); err != nil {
		fmt.Print(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user", handlers.CreateUserHandler(entClient))
	http.ListenAndServe(":10000", mux)

}
