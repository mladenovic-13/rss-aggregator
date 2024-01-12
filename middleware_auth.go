package main

import (
	"fmt"
	"net/http"

	"github.com/mladenovic-13/rss-aggregator/internal/auth"
	"github.com/mladenovic-13/rss-aggregator/internal/database"
)

type protectedHandler func(http.ResponseWriter, *http.Request, database.User)

func (ctx *Context) middlewareAuth(handler protectedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authorization error: %v", err))
			return
		}

		user, err := ctx.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("User not found: %v", err))
			return
		}

		handler(w, r, user)
	}
}
