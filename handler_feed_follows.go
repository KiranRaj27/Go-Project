package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Kiranraj27/go/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %v", err))
	}

	feedFollows, err := apiCg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create a feed %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFeedsFollows(feedFollows))
}

func (apiCg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create a feed %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedsToFeedsFollowsAll(feeds))
}

func (apiCg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	idString := chi.URLParam(r, "id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json %v", err))
	}
	err = apiCg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     id,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create a feed %v", err))
		return
	}
	respondWithJSON(w, 201, "Deleted")
}
