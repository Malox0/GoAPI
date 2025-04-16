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
)

type health struct {
	Status   string   `json:"status"`
	Messages []string `json:"messages"`
}

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("Error wiht DB connection: #{err.Error()}")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error with Driver #{err.Error()}")
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatalf("Error with migration #{err.Error()}")
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

	s := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
