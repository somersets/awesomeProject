package infrastructure

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occured while getting env variables, err:%s", err.Error())
	}
}
