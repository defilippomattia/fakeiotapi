package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() (*sql.DB, error) {

	connectionString := getConnString()

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func GetDB() *sql.DB {
	return db
}

func getConnString() string {
	host := getEnv("FRA_DB_HOST", "localhost")
	port := getEnv("FRA_DB_PORT", "6432")
	user := getEnv("FRA_DB_USER", "fakeiot")
	password := getEnv("FRA_DB_PASSWORD", "fakeiot")
	database := getEnv("FRA_DB_NAME", "fakeiot")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)
	return connectionString
}

// returns the value of the environment variable named by the key.
// If the variable is not present in the environment, the value fallback is returned.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
