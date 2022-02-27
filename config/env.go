package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	ELASTICSEARCH_URL string
}

const (
	ep string = ".env"
)

func InitEnv() {
	err := godotenv.Load(ep)

	if err != nil {
		log.Fatal("err loading .env")
	}
}

func GetConfig() Config {
	config := Config{
		ELASTICSEARCH_URL: "ELASTICSEARCH_URL",
	}

	return config
}
