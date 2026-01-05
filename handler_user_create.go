package main 

import (
"encoding/json"
"net/http"
"time"

"github.com/cahenrichs/Chirpy/internal/database"
"github.com/cahenrichs/Chirpy/internal/auth"

"github.com/google/uuid"
)

type User struct {
	ID	uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
	Email	string `json:"email"`
	Password string `json:"-"`
	IsChirpyRed bool `json:"is_chirpy_red"`
}

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
	}

	var params parameters

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters", err)
		return
	}

	if params.Email == ""{
		respondWithError(w, http.StatusBadRequest, "Email is required", nil)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Email: params.Email,
		HashedPassword: hashedPassword,
	})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
			return
		}
	

	respondWithJSON(w, http.StatusCreated, response{
		User: User{
			ID:	user.ID,
			CreatedAt:	user.CreatedAt,
			UpdatedAt:	user.UpdatedAt,
			Email:	user.Email,
			IsChirpyRed:	user.IsChirpyRed,
		},
	})
}