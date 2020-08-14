package handlers

import (
	"github.com/gorilla/mux"
)

func CreateRouter(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/login", login).Methods("POST")
	apiRouter.HandleFunc("/refresh", refresh).Methods("GET")

	moviesRouter := apiRouter.PathPrefix("/movies").Subrouter()
	moviesRouter.HandleFunc("/{movieId}", handleCreateAccountFavMovie).Methods("POST")
	moviesRouter.HandleFunc("/{movieId}", handleDeleteAccountFavMovie).Methods("DELETE")
}
