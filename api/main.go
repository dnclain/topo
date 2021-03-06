package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

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

	checkEnv()
}

// checkEnv ensures all necessary env data is present.
// panic in case of missing env.
// hide env vars with 'password', 'passwd' or 'key' in their name.
func checkEnv() {
	mandatoryEnvs := []string{
		ENV_POSTGRES_USER, ENV_POSTGRES_PASSWORD, ENV_POSTGRES_DB,
		ENV_POSTGRES_HOST, ENV_POSTGRES_PORT, ENV_API_PORT, ENV_POSTGRES_SCHEMA,
		ENV_MAX_FEATURE}

	reIsPassword := regexp.MustCompile(`(?i)password|passwd|key`)

	for _, theEnvName := range mandatoryEnvs {
		theEnv, isPresent := os.LookupEnv(theEnvName)
		if !isPresent {
			log.Fatalf("Sorry, env %v is mandatory. Please check environment variable or use --env <path> options\n", theEnvName)
		}

		if reIsPassword.MatchString(theEnvName) {
			theEnv = fmt.Sprintf("**(%v)**", len(theEnv))
		}

		fmt.Printf("* %s : %s \n", theEnvName, theEnv)
	}

	optionalEnvs := []string{ENV_VIEWER_URL, ENV_API_KEY}

	for _, theEnvName := range optionalEnvs {
		theEnv, isPresent := os.LookupEnv(theEnvName)

		if !isPresent {
			theEnv = "<undefined>"
		} else if reIsPassword.MatchString(theEnvName) {
			theEnv = fmt.Sprintf("**(%v)**", len(theEnv))
		}

		fmt.Printf("# %s : %s \n", theEnvName, theEnv)
	}

	// checks the format
	// TODO: replace by a more clever code.
	_maxFeatureRaw := os.Getenv(ENV_MAX_FEATURE)
	_maxFeature, err := strconv.Atoi(_maxFeatureRaw)
	if err != nil {
		log.Fatalf("%v should be a number but is '%v'\n", ENV_MAX_FEATURE, _maxFeatureRaw)
	}
	if _maxFeature < 0 {
		log.Fatalf("%v should be positive but is %v", ENV_MAX_FEATURE, _maxFeature)
	}
}

// main entry point.
func main() {

	fmt.Print("Database connection pool initialization : ")
	DB := initDB(os.Getenv(ENV_POSTGRES_USER),
		os.Getenv(ENV_POSTGRES_PASSWORD),
		os.Getenv(ENV_POSTGRES_DB),
		os.Getenv(ENV_POSTGRES_HOST),
		os.Getenv(ENV_POSTGRES_PORT))
	fmt.Println("✅")

	a := App{}
	a.Initialize(DB)
	a.Run(":" + os.Getenv("API_PORT"))
}

// initDB creates a connection pool from identifiers.
func initDB(user, password, dbname, hostname, port string) *sql.DB {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, hostname, port)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return DB
}
