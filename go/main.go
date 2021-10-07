package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// main entry point.
func main() {
	log.Println("Star")
	fmt.Print("Loading .env : ")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("✅")

	fmt.Print("Database connection pool initialization : ")
	initDB(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"))
	fmt.Println("✅")

	a := App{}
	a.Initialize()
	a.Run(":" + os.Getenv("APP_PORT"))
}

// initDB creates a connection pool from identifiers.
func initDB(user, password, dbname, hostname, port string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, hostname, port)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

}
