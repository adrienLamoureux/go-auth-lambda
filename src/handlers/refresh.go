package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type refreshResponse struct {
	Token      string `json:"token"`
	ExpireTime int64  `json:"expireTm"`
}

func refresh(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract the token from header
	tknStr := ""
	claim := &claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claim.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	refreshResponseJSON, err := json.Marshal(refreshResponse{
		Token:      tokenString,
		ExpireTime: expirationTime.Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(refreshResponseJSON)
}
