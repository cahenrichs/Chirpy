package main 

import (
	"encoding/json"
	"net/http"
	"errors"
	"database/sql"
	"github.com/cahenrichs/Chirpy/internal/auth"
	"log"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerWebhooks (w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data struct {
			UserID uuid.UUID `json:"user_id"`
		}
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}


log.Printf("API KEY HEADER: '%s'\n", apiKey)
log.Printf("POLKA KEY CFG:  '%s'\n", cfg.polkaKey)

	if apiKey != cfg.polkaKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couln't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = cfg.db.UpgradeToChirpyRed(r.Context(), params.Data.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
		respondWithError(w, http.StatusNotFound, "Couldn't find the user", err)
		return
	}
	respondWithError(w, http.StatusInternalServerError, "Couldn't update user", err)
	return
}

	w.WriteHeader(http.StatusNoContent)
}