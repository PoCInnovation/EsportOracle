package main

import (
	"esportoracle-backend/backend"
	"log"
	"net/http"
)

func main() {
	router := backend.SetupRoutes()
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
