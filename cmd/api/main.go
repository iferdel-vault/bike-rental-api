package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/iferdel-vault/bike-rental-api/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
)

type apiConfig struct {
	db        *database.Queries
	platform  string
	jwtSecret string
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	godotenv.Load(".env")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	dbConn, err := sql.Open("sqlite3", "./bikerental.db")
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	_ = apiConfig{
		db:        dbQueries,
		jwtSecret: jwtSecret,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/status", handlerStatus)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
