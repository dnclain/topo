package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Loading .env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"))
	a.Run(":" + os.Getenv("APP_PORT"))
}
