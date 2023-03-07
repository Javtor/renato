package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token string
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	strongestAvenger := goDotEnvVariable("STRONGEST_AVENGER")
	log.Println(strongestAvenger)
}
