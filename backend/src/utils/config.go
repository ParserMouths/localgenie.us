package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisClusterURL string
	RedisTLSDomain  string
	ServerPort      string
	PostgresURI     string
	JwtSecret       string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Print("Unable to load .env file. Continuing without loading it...")

	}

	config := &Config{
		RedisClusterURL: "redis://redis:6379",
		RedisTLSDomain:  "",
		ServerPort:      "6969",
		PostgresURI:     "postgresql://test:test@postgres/test",
		JwtSecret:       "",
	}

	envConfigVars := [...]string{
		"SERVER_PORT",
		"PG_DB_URI",
		"REDIS_CLUSTER_HOST",
		"JWT_SECRET",
	}

	for _, env := range envConfigVars {
		if os.Getenv(env) != "" {
			switch env {
			case "SERVER_PORT":
				config.ServerPort = os.Getenv(env)
			case "PG_DB_URI":
				config.PostgresURI = os.Getenv(env)
			case "REDIS_CLUSTER_HOST":
				config.RedisClusterURL = os.Getenv(env)
			case "JWT_SECRET":
				config.JwtSecret = os.Getenv(env)
			}
		}
	}

	return config
}
