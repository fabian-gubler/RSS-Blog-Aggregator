package main

import (
	"context"
	"net/http"

	_ "github.com/lib/pq"
)

func (cfg *apiConfig) handlerDeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	// Authentication and authorization checks should go here
	// For example, check if the request is from an admin user

	// Proceed to delete all users
	err := cfg.DB.DeleteAllUsers(context.Background())
	if err != nil {
		http.Error(w, "Failed to delete users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All users have been deleted."))
}
