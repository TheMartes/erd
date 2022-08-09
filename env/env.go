package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ep string = ".env"
)

func getValue(param string) string {
	err := godotenv.Load(ep)

	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(param)
}
