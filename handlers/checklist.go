// handlers/checklist.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"go-product-api/models"

	"github.com/gorilla/mux"
)

// CreateChecklist handles the creation of a new checklist.
func CreateChecklist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var checklist models.Checklist
		err := json.NewDecoder(r.Body).Decode(&checklist)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		username := r.Context().Value("username").(string)

		var userID int
		err = db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO checklists (user_id, item_name) VALUES (?, ?)",
			userID, checklist.ItemName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to get last insert id", http.StatusInternalServerError)
			return
		}
		checklist.ID = int(id)
		checklist.UserID = userID

		json.NewEncoder(w).Encode(checklist)
	}
}

// GetChecklists retrieves all checklists for a specific user.
func GetChecklists(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("username").(string)

		var userID int
		err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("SELECT id, user_id, item_name, created_at, updated_at FROM checklists WHERE user_id = ?", userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var checklists []models.Checklist
		for rows.Next() {
			var c models.Checklist
			var createdAt, updatedAt []uint8
			err = rows.Scan(&c.ID, &c.UserID, &c.ItemName, &createdAt, &updatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			c.CreatedAt, err = parseTime(createdAt)
			if err != nil {
				http.Error(w, "Error parsing created_at time", http.StatusInternalServerError)
				return
			}
			c.UpdatedAt, err = parseTime(updatedAt)
			if err != nil {
				http.Error(w, "Error parsing updated_at time", http.StatusInternalServerError)
				return
			}

			checklists = append(checklists, c)
		}

		json.NewEncoder(w).Encode(checklists)
	}
}

// GetChecklist retrieves a single checklist by its ID.
func GetChecklist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		checklistID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid checklist ID", http.StatusBadRequest)
			return
		}

		var checklist models.Checklist
		var createdAt, updatedAt []uint8
		err = db.QueryRow("SELECT id, user_id, item_name, created_at, updated_at FROM checklists WHERE id = ?", checklistID).Scan(&checklist.ID, &checklist.UserID, &checklist.ItemName, &createdAt, &updatedAt)
		if err != nil {
			http.Error(w, "Checklist not found", http.StatusNotFound)
			return
		}

		checklist.CreatedAt, err = parseTime(createdAt)
		if err != nil {
			http.Error(w, "Error parsing created_at time", http.StatusInternalServerError)
			return
		}
		checklist.UpdatedAt, err = parseTime(updatedAt)
		if err != nil {
			http.Error(w, "Error parsing updated_at time", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(checklist)
	}
}

// DeleteChecklist deletes a checklist by its ID.
func DeleteChecklist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		checklistID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid checklist ID", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("DELETE FROM checklists WHERE id = ?", checklistID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Checklist has Deleted"})
	}
}

// CreateItem handles the creation of a new item within a checklist.
func CreateItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		checklistID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid checklist ID", http.StatusBadRequest)
			return
		}

		var item models.Item
		err = json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO items (checklist_id, text) VALUES (?, ?)",
			checklistID, item.Text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to get last insert id", http.StatusInternalServerError)
			return
		}
		item.ID = int(id)
		item.ChecklistID = checklistID

		json.NewEncoder(w).Encode(item)
	}
}

// GetItems retrieves all items for a specific checklist.
func GetItems(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		checklistID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid checklist ID", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("SELECT id, checklist_id, text, completed, created_at, updated_at FROM items WHERE checklist_id = ?", checklistID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var items []models.Item
		for rows.Next() {
			var i models.Item
			var createdAt, updatedAt []uint8
			err = rows.Scan(&i.ID, &i.ChecklistID, &i.Text, &i.Completed, &createdAt, &updatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			i.CreatedAt, err = parseTime(createdAt)
			if err != nil {
				http.Error(w, "Error parsing created_at time", http.StatusInternalServerError)
				return
			}
			i.UpdatedAt, err = parseTime(updatedAt)
			if err != nil {
				http.Error(w, "Error parsing updated_at time", http.StatusInternalServerError)
				return
			}

			items = append(items, i)
		}

		json.NewEncoder(w).Encode(items)
	}
}

// GetItem retrieves a single item by its ID.
func GetItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		itemID, err := strconv.Atoi(vars["item_id"])
		if err != nil {
			http.Error(w, "Invalid item ID", http.StatusBadRequest)
			return
		}

		var item models.Item
		var createdAt, updatedAt []uint8
		err = db.QueryRow("SELECT id, checklist_id, text, completed, created_at, updated_at FROM items WHERE id = ?", itemID).Scan(&item.ID, &item.ChecklistID, &item.Text, &item.Completed, &createdAt, &updatedAt)
		if err != nil {
			http.Error(w, "Item not found", http.StatusNotFound)
			return
		}

		item.CreatedAt, err = parseTime(createdAt)
		if err != nil {
			http.Error(w, "Error parsing created_at time", http.StatusInternalServerError)
			return
		}
		item.UpdatedAt, err = parseTime(updatedAt)
		if err != nil {
			http.Error(w, "Error parsing updated_at time", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(item)
	}
}

// UpdateItem updates an existing item.
func UpdateItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		itemID, err := strconv.Atoi(vars["item_id"])
		if err != nil {
			http.Error(w, "Invalid item ID", http.StatusBadRequest)
			return
		}

		var item models.Item
		err = json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec("UPDATE items SET text = ?, completed = ? WHERE id = ?",
			item.Text, item.Completed, itemID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		item.ID = itemID
		json.NewEncoder(w).Encode(item)
	}
}

// DeleteItem deletes an item by its ID.
func DeleteItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		itemID, err := strconv.Atoi(vars["item_id"])
		if err != nil {
			http.Error(w, "Invalid item ID", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("DELETE FROM items WHERE id = ?", itemID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func parseTime(timeBytes []uint8) (time.Time, error) {
	timeString := string(timeBytes)
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}