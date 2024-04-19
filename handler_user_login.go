package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserWithToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Email        string `json:"email"`
	ID           int    `json:"id"`
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.GetUserByEmail(params.Email)
	if err != nil {
		respondWithError(w, 404, "Couldn't get user")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	signedAccess, err := generateAccessToken(user.ID, cfg.jwtSecret)
	signedRefresh, err := generateRefreshToken(user.ID, cfg.jwtSecret)

	if err != nil {
		respondWithError(w, 500, "Couldn't sign token")
		return
	}

	respondWithJSON(w, 200, UserWithToken{
		ID:           user.ID,
		Email:        user.Email,
		Token:        signedAccess,
		RefreshToken: signedRefresh,
	})

}
