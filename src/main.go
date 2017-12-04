package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	Success("> Listening at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
