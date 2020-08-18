package handlers

import (
	"github.com/gorilla/mux"
)

func CreateRouter(router *mux.Router) {
	router.HandleFunc("/ping", ping).Methods("GET")

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/login", login).Methods("POST")
	apiRouter.HandleFunc("/refresh", refresh).Methods("GET")

	apiRouter.HandleFunc("/movies", handleGetAccountFavMovies).Methods("GET")
	moviesRouter := apiRouter.PathPrefix("/movies").Subrouter()
	moviesRouter.HandleFunc("/{movieId}", handleCreateAccountFavMovie).Methods("POST")
	moviesRouter.HandleFunc("/{movieId}", handleDeleteAccountFavMovie).Methods("DELETE")
}
