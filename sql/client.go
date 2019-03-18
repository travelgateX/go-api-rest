package sql

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

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
func CreateClient() (*Client, error) {
	// Get database connection string
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(fmt.Errorf("fatal error config file: %s", err))
	}

	ssl := ""
	if !viper.GetBool("db.ssl") {
		ssl = "sslmode=disable"
	}

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %v",
		viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.user"), viper.GetString("db.pass"), viper.GetString("db.name"), ssl)

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
