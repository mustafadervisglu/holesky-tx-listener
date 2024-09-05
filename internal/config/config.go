package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Database string
	Ethereum string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return &Config{
		Database: os.Getenv("DB_CONNECTION"),
		Ethereum: os.Getenv("ETH_URL"),
	}
}
