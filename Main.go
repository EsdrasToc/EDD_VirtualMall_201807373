package main

import (
	"log"

	"net/http"

	"./Server"
)

func main() {
	go Server.TimerBlocks()
	s := Server.New()
	log.Fatal(http.ListenAndServe(":3000", s.Router()))
}
