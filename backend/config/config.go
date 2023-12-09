package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config func to get env value from key ---
func LocalEnvFile() {
	// Check if .env file exists
	_, err := os.Stat(".env")
	if err == nil {
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			fmt.Print("Error loading .env file")
		}
	}
}

func ReadEnvVariableWithDefault(key string) string {
	return ReadEnvVariable(key, true)
}

func ReadEnvVariable(key string, debug bool) string {
	envValue := os.Getenv(key)

	if envValue == "" {
		fmt.Printf("Error loading %s\n", key)
		return "" // or return some default value
	}

	if debug {
		fmt.Printf("%s: %s\n", key, envValue)
	}

	return envValue
}
