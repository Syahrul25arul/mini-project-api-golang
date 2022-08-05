package config

import (
	"fmt"
	"mini-project/logger"
	"os"
)

var SERVER_ADDRESS string
var SERVER_PORT string
var DB_USER string
var DB_PASSWORD string
var DB_HOST string
var DB_PORT string
var DB_NAME string
var TESTING string
var DB_NAME_TESTING string

func SanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"TESTING",
		"DB_NAME_TESTING",
	}

	for _, key := range envProps {
		if os.Getenv(key) == "" {
			logger.Fatal(fmt.Sprintf("environment variabel %s is not defined, application terminate", os.Getenv(key)))
		}
	}
	SERVER_ADDRESS = os.Getenv("SERVER_ADDRES")
	SERVER_PORT = os.Getenv("SERVER_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	TESTING = os.Getenv("TESTING")
	DB_NAME_TESTING = os.Getenv("DB_NAME_TESTING")
}
