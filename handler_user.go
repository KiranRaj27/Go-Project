package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Kiranraj27/go/internal/auth"
	"github.com/Kiranraj27/go/internal/database"
	"github.com/google/uuid"
)

func (apiCg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %v", err))
	}

	user, err := apiCg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create a user %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Couldn't find th api key %v", err))
		return
	}
	user, err := apiCg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
