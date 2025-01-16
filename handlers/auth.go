// handlers/auth.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"go-product-api/middleware"
	"go-product-api/models"

	"golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		_, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)",
			user.Username, string(hashedPassword), user.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		var storedUser models.User
		err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?",
			user.Username).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)

		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token, _ := middleware.GenerateToken(user.Username)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}
