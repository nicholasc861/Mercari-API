package main

import (
	"log"
	"net/http"
	"github.com/nicholasc861/mercari-api/app"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
