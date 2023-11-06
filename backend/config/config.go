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
func Config(key string) string {

	// Check if the environment variable exists and use it if available
	envValue := os.Getenv(key)
	if envValue == "" {
		fmt.Sprintf("Error loading %s",key)
	}
	return envValue
}
