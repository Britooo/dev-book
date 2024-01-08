package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// DefaultPort is the fallback port if none is specified in the environment.
const DefaultPort = 3000

// ConnectionString holds the connection string to the database.
// It should be initialized by calling LoadEnvironment.
var ConnectionString string

// Port holds the port number on which the API should run.
// It should be initialized by calling LoadEnvironment.
var Port int

// LoadEnvironment initializes configuration from environment variables.
func LoadEnvironment() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	initPort()

	initConnectionString()
}

// initPort reads the "API_PORT" environment variable and converts it to an integer.
// If the variable is not set or cannot be converted to an integer, it defaults to DefaultPort.
func initPort() {
	var err error
	portStr := os.Getenv("API_PORT")
	if portStr == "" {
		log.Printf("API_PORT not set, defaulting to %v", DefaultPort)
		Port = DefaultPort
		return
	}

	Port, err = strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Invalid API_PORT '%v', defaulting to %v", portStr, DefaultPort)
		Port = DefaultPort
	}
}

// initConnectionString constructs the database connection string from environment variables.
func initConnectionString() {
	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"),
	)
	if ConnectionString == "::@/?charset=utf8&parseTime=True&loc=Local" {
		log.Println("Warning: Database environment variables are not set. Connection string is incomplete.")
	}
}
