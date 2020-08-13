package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type loginRequestBody struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type loginResponse struct {
	Token      string `json:"token"`
	ExpireTime int64  `json:"expireTm"`
}

// Create the Signin handler
func Signin(w http.ResponseWriter, r *http.Request) {
	var creds loginRequestBody
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accEmailInfo, err := accountDatabase.GetAccountEmailByEmail(creds.Email)
	if err != nil || accEmailInfo == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	accInfo, err := accountDatabase.GetAccountInfo(accEmailInfo.AccID)
	if err != nil || accInfo == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if accInfo.Password != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claim := &claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	loginResponseJSON, err := json.Marshal(loginResponse{
		Token:      tokenString,
		ExpireTime: expirationTime.Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(loginResponseJSON)
}
