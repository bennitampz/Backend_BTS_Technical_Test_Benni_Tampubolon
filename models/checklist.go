// models/checklist.go
package models

import "time"

// Checklist represents a checklist item.
type Checklist struct {
	ID        int       `json:"id"`         // Unique identifier for the checklist
	UserID    int       `json:"user_id"`    // ID of the user who owns the checklist
	ItemName  string    `json:"item_name"`  // Name of the checklist
	CreatedAt time.Time `json:"created_at"` // Timestamp when the checklist was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the checklist was last updated
}

// Item represents an item within a checklist.
type Item struct {
	ID          int       `json:"id"`           // Unique identifier for the item
	ChecklistID int       `json:"checklist_id"` // ID of the checklist this item belongs to
	Text        string    `json:"text"`         // Text content of the item
	Completed   bool      `json:"completed"`    // Indicates if the item is completed
	CreatedAt   time.Time `json:"created_at"`   // Timestamp when the item was created
	UpdatedAt   time.Time `json:"updated_at"`   // Timestamp when the item was last updated
}
