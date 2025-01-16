// main.go
package main

import (
	"go-product-api/config"
	"go-product-api/handlers"
	"go-product-api/middleware"
	"go-product-api/migrations"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// Run migrations
	migrations.RunMigrations(db)

	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/register", handlers.Register(db)).Methods("POST")
	r.HandleFunc("/login", handlers.Login(db)).Methods("POST")

	// Protected checklist routes
	r.HandleFunc("/checklists", middleware.AuthMiddleware(handlers.GetChecklists(db))).Methods("GET")
	r.HandleFunc("/checklists", middleware.AuthMiddleware(handlers.CreateChecklist(db))).Methods("POST")
	r.HandleFunc("/checklists/{id}", middleware.AuthMiddleware(handlers.GetChecklist(db))).Methods("GET")
	r.HandleFunc("/checklists/{id}", middleware.AuthMiddleware(handlers.DeleteChecklist(db))).Methods("DELETE")

	// Protected item routes
	r.HandleFunc("/checklists/{id}/items", middleware.AuthMiddleware(handlers.GetItems(db))).Methods("GET")
	r.HandleFunc("/checklists/{id}/items", middleware.AuthMiddleware(handlers.CreateItem(db))).Methods("POST")
	r.HandleFunc("/checklists/{id}/items/{item_id}", middleware.AuthMiddleware(handlers.GetItem(db))).Methods("GET")
	r.HandleFunc("/checklists/{id}/items/{item_id}", middleware.AuthMiddleware(handlers.UpdateItem(db))).Methods("PUT")
	r.HandleFunc("/checklists/{id}/items/{item_id}", middleware.AuthMiddleware(handlers.DeleteItem(db))).Methods("DELETE")

	log.Println("Server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
