package pkg

import (
	"os"

	// Register some standard stuff
	_ "github.com/joho/godotenv/autoload"
)

// GetEnv ...
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}
