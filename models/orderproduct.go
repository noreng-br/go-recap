package models

// OrderProduct represents the 'order_product' junction table.
type OrderProduct struct {
    // Composite PK/FK 1
    OrderID   int `json:"order_id"` 
    // Composite PK/FK 2
    ProductID int `json:"product_id"` 

    Quantity int     `json:"quantity"`
    UnitPrice float64 `json:"unit_price"` 
}
