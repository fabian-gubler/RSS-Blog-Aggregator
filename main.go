package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	serverHits int
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	apiCfg := apiConfig{
		serverHits: 0,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("GET /v1/readiness", apiCfg.handlerReadiness)
	mux.HandleFunc("GET /v1/err", apiCfg.handlerErr)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
