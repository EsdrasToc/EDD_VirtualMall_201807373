package main

import (
	"log"

	"net/http"

	"./Server"
)

func main() {
	s := Server.New()
	log.Fatal(http.ListenAndServe(":3000", s.Router()))
}
