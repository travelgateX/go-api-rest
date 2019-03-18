package main

import (
	"billing-calculation-center/sql"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := defaultPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	sql.MustCreateClient()

	router := newRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
