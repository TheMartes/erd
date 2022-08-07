package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ep string = ".env"
)

func GetEnvValue(param string) string {
	err := godotenv.Load(ep)

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(param)
}
