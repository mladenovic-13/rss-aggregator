package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mladenovic-13/rss-aggregator/internal/database"
)

func (ctx *Context) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Data struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	data := Data{}
	err := decoder.Decode(&data)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := ctx.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      data.Name,
		Url:       data.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (ctx *Context) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := ctx.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to get feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
