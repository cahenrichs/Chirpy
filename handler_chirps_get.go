package main

import (
	"net/http"

	"github.com/google/uuid"
)


func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chirp ID", err)
		return
	}

	dbChirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:		dbChirp.ID,
		CreatedAt:	dbChirp.CreatedAt,
		UpdatedAt:	dbChirp.UpdatedAt,
		Body:		dbChirp.Body,
		UserID:		dbChirp.UserID,
	})
}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not get Chirps", err)
		return
	}


	authorID := uuid.Nil
authorIDString := r.URL.Query().Get("author_id")
if authorIDString != "" {
    authorID, err = uuid.Parse(authorIDString)
    if err != nil {
        // 400
        return
    }
}

chirps := []Chirp{}
for _, dbChirp := range dbChirps {
    if authorID != uuid.Nil && dbChirp.UserID != authorID {
        continue
    }
    
	chirps = append(chirps, Chirp{
        ID:        dbChirp.ID,
        CreatedAt: dbChirp.CreatedAt,
        UpdatedAt: dbChirp.UpdatedAt,
        UserID:    dbChirp.UserID,
        Body:      dbChirp.Body,
    })
	
	respondWithJSON(w, http.StatusOK, chirps)

	}
}

/*
	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:		dbChirp.ID,
			CreatedAt:	dbChirp.CreatedAt,
			UpdatedAt:	dbChirp.UpdatedAt,
			Body:		dbChirp.Body,
			UserID:		dbChirp.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)
}
*/