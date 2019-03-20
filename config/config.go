package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config struct
type Config struct{}

// LoadConfig load configuration file
func LoadConfig() Config {
	// Get database connection string
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(fmt.Errorf("fatal error config file: %s", err))
	}

	return Config{}
}

// GetSQLConnectionString get sql connection string
func (c *Config) GetSQLConnectionString() string {
	ssl := ""
	if !viper.GetBool("db.ssl") {
		ssl = "sslmode=disable"
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %v",
		viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.user"), viper.GetString("db.pass"), viper.GetString("db.name"), ssl)
}
