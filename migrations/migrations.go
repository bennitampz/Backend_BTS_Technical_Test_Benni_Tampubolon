// migrations/migrations.go
package migrations

import (
	"database/sql"
	"log"
)

// RunMigrations executes all database migrations.
// It takes a *sql.DB object as input, which represents the database connection.
func RunMigrations(db *sql.DB) {
	// Call the functions to create the users, checklists, and items tables.
	createUsersTable(db)
	createChecklistsTable(db)
	createItemsTable(db)
	// Log a message indicating that the migrations have completed successfully.
	log.Println("Migrations completed successfully")
}

// createUsersTable creates the users table if it doesn't exist.
func createUsersTable(db *sql.DB) {
	// SQL query to create the users table.
	// The table includes columns for id, username, password, email, and created_at.
	// The username and email columns are unique and not null.
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

	// Execute the SQL query.
	_, err := db.Exec(query)
	if err != nil {
		// If there's an error during table creation, log a fatal error.
		log.Fatal("Error creating users table:", err)
	}
}

// createChecklistsTable creates the checklists table if it doesn't exist.
func createChecklistsTable(db *sql.DB) {
	// SQL query to create the checklists table.
	// The table includes columns for id, user_id, item_name, created_at, and updated_at.
	// The user_id column is a foreign key referencing the users table.
	query := `
    CREATE TABLE IF NOT EXISTS checklists (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT NOT NULL,
        item_name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id)
    )`

	// Execute the SQL query.
	_, err := db.Exec(query)
	if err != nil {
		// If there's an error during table creation, log a fatal error.
		log.Fatal("Error creating checklists table:", err)
	}
}

// createItemsTable creates the items table if it doesn't exist.
func createItemsTable(db *sql.DB) {
	// SQL query to create the items table.
	// The table includes columns for id, checklist_id, text, completed, created_at, and updated_at.
	// The checklist_id column is a foreign key referencing the checklists table.
	query := `
    CREATE TABLE IF NOT EXISTS items (
        id INT AUTO_INCREMENT PRIMARY KEY,
        checklist_id INT NOT NULL,
        text VARCHAR(255) NOT NULL,
        completed BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        FOREIGN KEY (checklist_id) REFERENCES checklists(id)
    )`

	// Execute the SQL query.
	_, err := db.Exec(query)
	if err != nil {
		// If there's an error during table creation, log a fatal error.
		log.Fatal("Error creating items table:", err)
	}
}
