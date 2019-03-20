package main

import (
	"log"
	"net/http"
	"os"

	"go-api-rest/sql"
)

const defaultPort = "8080"

func main() {
	port := defaultPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// Create client
	sql.MustCreateClient()

	router := newRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
