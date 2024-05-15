package main

import "net/http"

type Status struct {
	Status string `json:"status"`
}

type Error struct {
	Error string `json:"error"`
}

func (cfg *apiConfig) handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, Status{
		Status: "ok",
	})
}

func (cfg *apiConfig) handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server error")
}
