package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("./env/local.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
		panic(err)
	}
}
