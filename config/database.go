package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB establishes a connection to the MySQL database.
// It returns a *sql.DB object, which represents the database connection pool.
func ConnectDB() (*sql.DB, error) {
	// Open a connection to the MySQL database.
	// The connection string specifies the database driver (mysql), the user (root),
	// the host (localhost:3306), and the database name (noteApp_db).
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/noteApp_db")
	if err != nil {
		// If there's an error during connection, return nil and the error.
		return nil, err
	}
	// If the connection is successful, return the database connection pool and nil error.
	return db, nil
}
