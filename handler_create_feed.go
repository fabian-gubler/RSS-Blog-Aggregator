package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fabian-gubler/RSSFlow/internal/database"

	"github.com/google/uuid"
)

type FeedParams struct {
	Name string
	URL  string
}

type FeedResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"created_at"`
	UpdateAt  string    `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Generate UUID for User
	newUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return
	}

	// TODO: This is a temporary solution to fake a User ID
	// Generate UUID for User
	newUserUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return
	}

	// Prepare parameters for the database function
	dbParams := database.CreateFeedParams{
		Name:      params.Name,
		ID:        newUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       params.URL,
		UserID:    newUserUUID, // TODO: make this work automatically with authentication
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), dbParams)
	if err != nil {
		// print error received from database
		fmt.Printf("Error creating feed: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed: %v")
		return
	}

	respondWithJSON(w, http.StatusCreated, FeedResponse{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt.Format(time.RFC3339),
		UpdateAt:  feed.UpdatedAt.Format(time.RFC3339),
		Name:      feed.Name,
		UserID:    newUUID,
	})
}
