// +build !lambda

package main

import (
	"log"
	"net/http"

	"github.com/adrienLamoureux/go-auth-lambda/src/handlers"
	"github.com/gorilla/mux"
)

func createDevRouter() *mux.Router {
	router := mux.NewRouter()
	handlers.CreateRouter(router)
	return router
}

func main() {
	log.Fatal(http.ListenAndServe(":7200", createDevRouter()))
}
