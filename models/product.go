package models

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CategoryID  string `json:"category_id"` // Foreign Key relationship
	Description string `json:"description"`
}
