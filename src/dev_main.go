// +build !lambda

package main

import (
	"log"
	"net/http"

	"github.com/adrienLamoureux/go-auth-lambda/src/databases/dynamodb"
	"github.com/adrienLamoureux/go-auth-lambda/src/handlers"
	"github.com/gorilla/mux"
)

func createDevRouter() *mux.Router {
	router := mux.NewRouter()
	handlers.CreateRouter(router)
	devRouter := router.PathPrefix("/dev").Subrouter()
	devRouter.HandleFunc("/createDynamoTables", createDynamoTables).Methods("POST")
	return router
}

func main() {
	log.Fatal(http.ListenAndServe(":7200", createDevRouter()))
}

func createDynamoTables(w http.ResponseWriter, r *http.Request) {
	err := dynamodb.CreateAccountTable()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	err = dynamodb.CreateAccountEmailTable()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	err = dynamodb.CreateAccountFavMovieTable()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
}
