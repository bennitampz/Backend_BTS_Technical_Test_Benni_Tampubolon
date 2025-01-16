package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ConnectDB establishes a connection to the MySQL database.
// It returns a *sql.DB object, which represents the database connection pool.
func ConnectDB() (*sql.DB, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database URL from environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Open a connection to the MySQL database.
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		// If there's an error during connection, return nil and the error.
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// If the connection is successful, return the database connection pool and nil error.
	return db, nil
}
