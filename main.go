package main

import (
	"log"
	"net/http"
	"os"

	"go-api-rest/config"
	"go-api-rest/sql"
)

const defaultPort = "8080"

func main() {
	port := defaultPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// Load config
	c := config.LoadConfig()

	// Create client
	sql.MustCreateClient(c.GetSQLConnectionString())

	// Create new router
	router := newRouter(c.GetJwtParserConfig())

	log.Fatal(http.ListenAndServe(":"+port, router))
}
