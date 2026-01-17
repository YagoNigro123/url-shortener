package main

import (
	"context"
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
		log.Println("No .env file found, assuming variables are set in environment")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = os.Getenv("DB_URL")
	}

	if dbURL == "" {
		log.Fatal("Error: DATABASE_URL (or DB_URL) environment variable is not set")
	}

	postgresStore, err := store.NewPostgresStore(dbURL)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	defer postgresStore.Close()

	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	redisPassword := os.Getenv("REDISPASSWORD")

	var redisAddr string

	if redisHost != "" && redisPort != "" {
		redisAddr = fmt.Sprintf("%s:%s", redisHost, redisPort)
	} else {
		redisAddr = os.Getenv("REDIS_ADDR")
		if redisAddr == "" {
			redisAddr = "localhost:6379"
		}
	}

	redisClient := store.NewRedisClient(redisAddr, redisPassword)

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Printf("Warning: Could not connect to Redis at %s: %v", redisAddr, err)
	} else {
		fmt.Println("Connected to Redis successfully")
	}

	svc := core.NewService(postgresStore, redisClient)
	handler := api.NewHandler(svc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/shorten", handler.CreateLink)
	r.Get("/{id}", handler.Redirect)

	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public", fileServer))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
