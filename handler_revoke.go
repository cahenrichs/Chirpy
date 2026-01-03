package main

import (
"github.com/cahenrichs/Chirpy/internal/auth"
"net/http"
)

func(cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't get token", err)
		return
	}

	_, err = cfg.db.RevokeRefreshToken(r.Context(), refreshToken) 
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke session", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}