// +build !lambda

package main

import (
	"log"
	"net/http"

	"github.com/adrienLamoureux/go-auth-lambda/src/handlers"
)

func main() {
	log.Fatal(http.ListenAndServe(":7200", handlers.CreateRouter()))
}
