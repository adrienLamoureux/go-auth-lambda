package handlers

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func getClaimsFromRequest(r *http.Request) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}
	tkn, err := jwt.ParseWithClaims(r.Header.Get("Authorization"), claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("Invalid token")
	}
	return claims, nil
}
