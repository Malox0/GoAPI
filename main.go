package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/gorilla/mux"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

type health struct {
	Status   string   `json:"status"`
	Messages []string `json:"messages"`
}

type jsonError struct {
	Code string `json: code`
	Msg  string `json: message`
}

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("Error wiht DB connection: %v", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error with Driver %v", err.Error())
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://migrations/",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatalf("Error with migration %v", err.Error())
	}

	migrator.Steps(2)

	fmt.Println("Please make Docker work")
	fmt.Println("Pls speed up docker")
	r := mux.NewRouter()

	r.HandleFunc(
		"/healthcheck",
		func(w http.ResponseWriter, r *http.Request) {
			h := health{
				Status:   "OK",
				Messages: []string{},
			}

			b, _ := json.Marshal(h)

			w.Write(b)
			w.WriteHeader(http.StatusOK)
		})

	r.HandleFunc(
		"/book/{isbn}",
		func(w http.ResponseWriter, r *http.Request) {

			v := mux.Vars(r)

			e := jsonError{
				Code: "001",
				Msg:  fmt.Sprintf("No book with ISBN %v", v["isbn"]),
			}

			b, _ := json.Marshal(e)

			w.WriteHeader(http.StatusNotFound)
			w.Write(b)

		})

	s := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
