package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	serverHits := cfg.serverHits
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API has been used " + fmt.Sprintf("%d", serverHits) + " times!"))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.serverHits++
		next.ServeHTTP(w, r)
	})
}
