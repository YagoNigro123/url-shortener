package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YagoNigro123/url-shortener/internal/api"
	"github.com/YagoNigro123/url-shortener/internal/core"
	"github.com/YagoNigro123/url-shortener/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	connStr := "postgres://user:password@localhost:5432/shortener_db?sslmode=disable"

	postgressStore, err := store.NewPostgresStore(connStr)
	if err != nil {
		log.Fatalf("error with conect to db: %v", err)
	}

	//defer postgresStore().Close()

	svc := core.NewService(postgressStore)

	handler := api.NewHandler(svc)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/shorten", handler.CreateLink)
	r.Get("/{id}", handler.Redirect)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed: %v", err)
	}
}
