// +build !lambda

package main

import (
	"fmt"
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
	port := "7200"
	fmt.Printf("Starting server at %s\n", port)
	err := http.ListenAndServe(":"+port, createDevRouter())
	if err != nil {
		fmt.Printf("Failed to start server at port %s\nError: %s", port, err.Error())
		panic(err)
	}
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
