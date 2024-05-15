package main

import (
	"context"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {

	apiKey := r.Header.Get("Authorization")

	user, err := cfg.DB.GetUserByAPIKey(context.Background(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't find user: %v")
		return
	}

	respondWithJSON(w, http.StatusCreated, UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdateAt:  user.UpdatedAt.Format(time.RFC3339),
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	})
}


