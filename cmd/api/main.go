package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/YagoNigro123/url-shortener/internal/api"
	"github.com/YagoNigro123/url-shortener/internal/core"
	"github.com/YagoNigro123/url-shortener/internal/store"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No . env file found, assuming variables are set in evironment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	postgresStore, err := store.NewPostgresStore(dbURL)
	if err != nil {
		log.Fatalf("error with conect to db: %v", err)
	}

	defer postgresStore.Close()

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisClient := store.NewRedisClient(redisAddr, "")

	svc := core.NewService(postgresStore, redisClient)

	handler := api.NewHandler(svc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/shorten", handler.CreateLink)
	r.Get("/{id}", handler.Redirect)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed: %v", err)
	}
}
