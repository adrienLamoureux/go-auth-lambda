package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleCreateAccountFavMovie(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract the token from header
	accID := "a"
	//movieID := strings.TrimPrefix(r.URL.Path, "/movies/")
	params := mux.Vars(r)
	movieID := params["movieId"]
	if len(movieID) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err := accountDatabase.CreateAccountFavMovie(accID, movieID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}

func handleDeleteAccountFavMovie(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract the token from header
	accID := "a"
	//movieID := strings.TrimPrefix(r.URL.Path, "/movies/")
	params := mux.Vars(r)
	movieID := params["movieId"]
	if len(movieID) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err := accountDatabase.DeleteAccountFavMovie(accID, movieID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}
