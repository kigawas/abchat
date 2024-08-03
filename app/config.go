package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host        string
	Port        string
	DatabaseURL string
	RedisURL    string
	Prefork     bool
}

func FromEnv() Config {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")
	redisUrl := os.Getenv("REDIS_URL")
	prefork := os.Getenv("PREFORK")

	return Config{
		Host:        host,
		Port:        port,
		DatabaseURL: dbUrl,
		RedisURL:    redisUrl,
		Prefork:     prefork == "1",
	}
}

func (c Config) URL() string {
	return c.Host + ":" + c.Port
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
