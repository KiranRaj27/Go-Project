package main

import (
	"fmt"
	"net/http"

	"github.com/Kiranraj27/go/internal/auth"
	"github.com/Kiranraj27/go/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	}
}
