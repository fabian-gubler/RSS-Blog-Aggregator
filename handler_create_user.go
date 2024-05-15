package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fabian-gubler/RSSFlow/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type UserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"created_at"`
	UpdateAt  string    `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Check if User exists
	userChecked, err := cfg.DB.GetUser(context.Background(), params.Name)
	if userChecked != (database.User{}) {
		respondWithError(w, http.StatusBadRequest, "User already exists")
		return
	}

	// Generate UUID for User
	newUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return
	}

	// Prepare parameters for the database function
	dbParams := database.CreateUserParams{
		ID:        newUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	}

	// Create user
	user, err := cfg.DB.CreateUser(context.Background(), dbParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user: %v")
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
