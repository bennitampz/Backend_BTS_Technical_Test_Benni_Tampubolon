package models

// User represents a user in the system.
type User struct {
	ID       int    `json:"id"`       // Unique identifier for the user
	Username string `json:"username"` // Username of the user
	Password string `json:"password"` // Password of the user (hashed)
	Email    string `json:"email"`    // Email address of the user
}
