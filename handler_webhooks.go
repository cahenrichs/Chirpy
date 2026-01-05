package main 

import (
	"encoding/json"
	"net/http"
	"errors"
	"database/sql"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerWebhooks (w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data struct {
			UserID uuid.UUID `json:"user_id"`
		}
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
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