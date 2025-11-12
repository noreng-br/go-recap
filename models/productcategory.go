package models

// ProductCategory represents the 'product_category' junction table.
type ProductCategory struct {
    // Composite PK/FK 1
    ProductID  int `json:"product_id"`
    // Composite PK/FK 2
    CategoryID int `json:"category_id"`
}
