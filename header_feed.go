package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/habibbushira/rssaggregator/internal/database"

	"github.com/google/uuid"
)

func (apiCfg apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `name`
		Url  string `url`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create feed %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeed(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get feeds %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
