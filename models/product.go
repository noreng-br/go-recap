package models

type Product struct {
	ProductID          string `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
  Price float64 `json:"price"`
}

type ProductDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
  Price float64 `json:"price"`
}
