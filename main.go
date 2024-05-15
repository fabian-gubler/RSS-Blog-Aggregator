package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/fabian-gubler/RSSFlow/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	serverHits int
	DB         *database.Queries
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %v\n", err)
	}

	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error connecting to database: %v\n", err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		serverHits: 0,
		DB:         dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("GET /v1/readiness", apiCfg.handlerReadiness)
	mux.HandleFunc("GET /v1/err", apiCfg.handlerErr)

	mux.HandleFunc("GET /v1/users", apiCfg.handlerGetUser)
	mux.HandleFunc("POST /v1/users", apiCfg.handlerCreateUser)
	mux.HandleFunc("DELETE /v1/users/all", apiCfg.handlerDeleteAllUsers)

	mux.HandleFunc("POST /v1/feeds", apiCfg.handlerCreateFeed)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
