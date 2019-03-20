package sql

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"

	// Required
	_ "github.com/lib/pq"
)

// Instance : global variable for client
var Instance *Client

var once sync.Once

// Client : sqlx client
type Client struct {
	DB *sqlx.DB
}

// CreateClient : new sqlx client constructor
func CreateClient(connString string) (*Client, error) {
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("DBConnect: Fail. ", err)
	}

	c := &Client{DB: db}
	return c, err
}

// MustCreateClient : singleton pattern for new sqlx client
func MustCreateClient(connString string) {
	once.Do(func() {
		var err error
		Instance, err = CreateClient(connString)
		if err != nil {
			log.Fatal(err)
		}
	})
}
