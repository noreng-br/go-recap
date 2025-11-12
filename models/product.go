package models

type Product struct {
	ProductID          string `json:"product_id"`
	Name        string `json:"name"`
	CategoryID  string `json:"category_id"` // Foreign Key relationship
	Description string `json:"description"`
  Price float64 `json:"price"`
}
