package handlers

import (
	"github.com/gorilla/mux"
)

func CreateRouter(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/login", Signin).Methods("POST")
	apiRouter.HandleFunc("/refresh", Signin).Methods("GET")
}
