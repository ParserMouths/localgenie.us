package utils

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisClusterURL string
	RedisTLSDomain  string
	ServerPort      int32
	PostgresURI     string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Print("Unable to load .env file. Continuing without loading it...")

	}

	config := &Config{
		RedisClusterURL: "redis://localhost:6379",
		RedisTLSDomain:  "",
		ServerPort:      6969,
		PostgresURI:     "postgresql://test:test@localhost/test",
	}

	return config
}
