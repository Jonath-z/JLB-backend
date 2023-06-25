package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	err := godotenv.Load()
	
	if err != nil{
		log.Fatal("Cant not load environment variables")
	}
}