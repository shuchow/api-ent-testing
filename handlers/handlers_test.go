package handlers

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuchow/api-ent-testing/ent"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	entClient := getMockDBClient()
	defer entClient.Close()

	w := getResponseWriter()

	// Mock request object with dummy data.
	r := http.Request{}
	r.Header = http.Header{"Content-Type": []string{"application/json"}, "Accept": []string{"application/json"}}
	r.Body = io.NopCloser(strings.NewReader(
		`{"userName": "test",
			"password": "user", 
			"email": "nobody@localhost.com"}`))

	// get the http.HandlerFunc directly from the API handler and call it.
	createUserHandler := CreateUserHandler(entClient)
	createUserHandler(w, &r)

	if w.Code != http.StatusCreated {
		t.Errorf("CreateUserHandler returned wrong status code: got %v want %v", w.Code, http.StatusCreated)
	}

	clientResults, err := entClient.User.Query().All(context.Background())
	if err != nil {
		t.Errorf(err.Error())
	} else {
		if len(clientResults) != 1 {
			t.Errorf("More than one results returned.  Got: %d", len(clientResults))
		}

		if clientResults[0].UserName != "test" {
			t.Errorf("Wrong username returned.  Got: %s", clientResults[0].UserName)
		}

		if clientResults[0].Email != "nobody@localhost.com" {
			t.Errorf("Wrong email returned.  Got: %s", clientResults[0].Email)
		}
	}
}

func TestBadPayload(t *testing.T) {
	entClient := getMockDBClient()
	defer entClient.Close()

	// Mock request object with dummy data.
	r := http.Request{}
	r.Header = http.Header{"Content-Type": []string{"application/json"}, "Accept": []string{"application/json"}}
	r.Body = io.NopCloser(strings.NewReader(
		`{"userName": "test",`))

	w := getResponseWriter()

	// get the http.HandlerFunc directly from the API handler and call it.
	createUserHandler := CreateUserHandler(entClient)
	createUserHandler(w, &r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("CreateUserHandler returned wrong status code: got %v want %v", w.Code, http.StatusBadRequest)
	}
}

func getResponseWriter() *httptest.ResponseRecorder {
	rw := httptest.NewRecorder()
	return rw
}

func getMockDBClient() *ent.Client {
	// Create a DBClient using in memory SQL Lite.  Nothing persisted.

	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %#v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Print(err)
	}

	return client

}
