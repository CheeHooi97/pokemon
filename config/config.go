package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Env           string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	SystemAesKey  string
	PokemonApiKey string
)

// LoadConfig
func LoadConfig() {
	_ = godotenv.Load()

	Env = GetEnv("ENV")
	DBHost = GetEnv("MYSQL_HOST")
	DBPort = GetEnv("MYSQL_PORT")
	DBUser = GetEnv("MYSQL_USER")
	DBPassword = GetEnv("MYSQL_PASSWORD")
	DBName = GetEnv("MYSQL_DATABASE")
	SystemAesKey = GetEnv("SYSTEM_AES_KEY")
	PokemonApiKey = os.Getenv("POKEMON_API_KEY") // Optional, so use os.Getenv directly
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("%s environment variable not set", key)
	}
	return value
}
