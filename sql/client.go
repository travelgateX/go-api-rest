package sql

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"

	// Required
	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = "5432"
	User     = "postgres"
	Password = "postgres"
	Dbname   = "postgres"
)

// Instance : global variable for client
var Instance *Client

var once sync.Once

// Client : sqlx client
type Client struct {
	DB *sqlx.DB
}

// CreateClient : new sqlx client constructor
func CreateClient() (*Client, error) {
	ssl := "sslmode=disable"

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %v",
		Host, Port, User, Password, Dbname, ssl)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("DBConnect: Fail. ", err)
	}

	c := &Client{DB: db}
	return c, err
}


// MustCreateClient : singleton pattern for new sqlx client
func MustCreateClient() {
	once.Do(func() {
		var err error
		Instance, err = CreateClient()
		if err != nil {
			log.Fatal(err)
		}
	})
}
