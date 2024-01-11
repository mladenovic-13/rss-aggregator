package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mladenovic-13/rss-aggregator/internal/database"
)

func (ctx *Context) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	data := Data{}
	err := decoder.Decode(&data)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := ctx.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      data.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
