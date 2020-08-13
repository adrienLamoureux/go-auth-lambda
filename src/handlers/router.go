package handlers

import (
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/login", Signin).Methods("POST")
	apiRouter.HandleFunc("/welcome", Signin).Methods("GET")
	apiRouter.HandleFunc("/refresh", Signin).Methods("GET")
	return apiRouter
}
