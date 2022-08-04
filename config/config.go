package config

import (
	"fmt"
	"mini-project/logger"
	"os"
)

var SERVER_ADDRESS string
var SERVER_PORT string

func SanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}

	for _, key := range envProps {
		if os.Getenv(key) == "" {
			logger.Fatal(fmt.Sprintf("environment variabel %s is not defined, application terminate", os.Getenv(key)))
		}
	}
	SERVER_ADDRESS = os.Getenv("SERVER_ADDRES")
	SERVER_PORT = os.Getenv("SERVER_PORT")
}
