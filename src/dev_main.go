// +build !lambda

package main

import (
	"log"
	"net/http"

	"github.com/adrienLamoureux/go-auth-lambda/src/handlers"
)

func main() {
	handlers.CreateRouter()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
