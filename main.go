package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Println("API is Read")
	log.Fatal(http.ListenAndServe(":8080", router))
}
