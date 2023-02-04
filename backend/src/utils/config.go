package utils

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisClusterURL       string
	RedisTLSDomain        string
	ServerPort            string
	PostgresURI           string
	JwtSecret             string
	StoryBlokOAuth        string
	NotificationPublicKey string
	SpaceID               string
	ProjectRoot           string
	NotifPrivateKey       string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Print("Unable to load .env file. Continuing without loading it...")

	}

	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))

	config := &Config{
		RedisClusterURL:       "redis://redis:6379",
		RedisTLSDomain:        "",
		ServerPort:            "6969",
		PostgresURI:           "postgresql://test:test@postgres/test",
		JwtSecret:             "",
		StoryBlokOAuth:        "",
		NotificationPublicKey: "BMLTD4SXRjPwfFAWZCOcv9_IyWoMGr1FX1SLTgtMdTLkh5NJu6qODaju484eyptfd1m7IZl037nDQMXPcfMpRUE",
		SpaceID:               "195405",
		ProjectRoot:           filepath.Dir(d),
		NotifPrivateKey:       "",
	}

	envConfigVars := [...]string{
		"SERVER_PORT",
		"PG_DB_URI",
		"REDIS_CLUSTER_HOST",
		"JWT_SECRET",
		"STORYBLOK_OAUTH",
		"NOTIF_PRIVATE_KEY",
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
			case "STORYBLOK_OAUTH":
				config.StoryBlokOAuth = os.Getenv(env)
			case "NOTIF_PRIVATE_KEY":
				config.NotifPrivateKey = os.Getenv(env)
			}
		}
	}

	return config
}
