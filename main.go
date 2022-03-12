package main

import (
	"log"
	"net/http"

	"github.com/ivanauliaa/simple-rest-mux/router"
)

func main() {
	router := router.Router()

	log.Fatal(http.ListenAndServe(":8081", router))
}
