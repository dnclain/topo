package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("Initializing...")

	_flagEnv := flag.String("env", "", "Select a .env file for specific configuration and env variable overload.")

	flag.Parse()

	if *_flagEnv != "" {
		fmt.Printf("Loading .env : %s", *_flagEnv)
		err := godotenv.Load(*_flagEnv)
		if err != nil {
			log.Fatal("Error loading .env file : ", *_flagEnv)
		}
		fmt.Println("✅")
	}

}

// main entry point.
func main() {

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
