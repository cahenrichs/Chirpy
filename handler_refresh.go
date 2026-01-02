package main

import (
"github.com/cahenrichs/Chirpy/internal/auth"
"net/http"
"time"

)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {

	type response struct {
		Token string `json:"token"`
	}

	//Get the token
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find token", err)
		return
	}

	//look up the refresh token
	user, err := cfg.db.GetUserFromRefreashToken(r.Context(), refreshToken) 
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get user for refreash token", err)
		return
	}

	accessToken, err := auth.MakeJWT(
		user.ID,
		cfg.jwtSecret,
		time.Hour,
	)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Token:accessToken,
	})

}