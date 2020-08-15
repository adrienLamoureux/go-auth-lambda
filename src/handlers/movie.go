package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adrienLamoureux/go-auth-lambda/src/databases"
	"github.com/gorilla/mux"
)

type getAccountFavMoviesResponse struct {
	AccountFavMovies []*getAccountFavMovieResponse `json:"accountFavMovies"`
}

type getAccountFavMovieResponse struct {
	MovieID string `json:"movieId"`
}

func handleGetAccountFavMovies(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaimsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	accID := claims.Id
	accountFavMovieInfoList, err := accountDatabase.GetAccountFavMoviesInfo(accID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	accountFavMovieList := make([]*getAccountFavMovieResponse, len(accountFavMovieInfoList))
	for i, accountFavMovieInfo := range accountFavMovieInfoList {
		accountFavMovieList[i] = &getAccountFavMovieResponse{
			MovieID: accountFavMovieInfo.MovieID,
		}
	}
	accountFavMoviesResponse, err := json.Marshal(getAccountFavMoviesResponse{
		AccountFavMovies: accountFavMovieList,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(accountFavMoviesResponse)
}

func handleCreateAccountFavMovie(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaimsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	accID := claims.Id
	params := mux.Vars(r)
	movieID := params["movieId"]
	if len(movieID) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err = accountDatabase.CreateAccountFavMovie(&databases.AccountFavMovieInfo{
		AccID:   accID,
		MovieID: movieID,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}

func handleDeleteAccountFavMovie(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaimsFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	accID := claims.Id
	params := mux.Vars(r)
	movieID := params["movieId"]
	if len(movieID) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err = accountDatabase.DeleteAccountFavMovie(accID, movieID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}
