package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChripValidate(w http.ResponseWriter, r *http.Request) {
	type chirpParams struct {
		Body string `json:"body"`
	}

	type validResponce struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := chirpParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert": {},
		"fornax": {},
	}

	cleaned := getCleanedBody(params.Body, badWords)

	respondWithJSON(w, http.StatusOK, validResponce{
		CleanedBody: cleaned,
	})
}

func getCleanedBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")

	for i, word := range words {
		word = strings.ToLower(word)
		if _, ok := badWords[word]; ok {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}