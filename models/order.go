package models

import (
  "time"
)

type Order struct {
  OrderID  string  `json:"order_id"`
  UserId  string `json:"user_id"`
  OrderedDate time.Time  `json:"ordered_date"`
  DeliverDate time.Time  `json:"deliver_date"`
  Products []Product `json:"products"`
  Status string `json:"status"`
}

// OrderRequest struct represents the incoming API data for a new order.
type OrderRequest struct {
    UserID int `json:"user_id"`
    Status string `json:"status"`
    // Array of products to be included in the order
    Items []OrderItemRequest `json:"items"`
}

// OrderItemRequest captures the details needed for one product in the order.
type OrderItemRequest struct {
    ProductID int `json:"product_id"`
    Quantity  int `json:"quantity"`
    // Note: We'll retrieve UnitPrice from the products table for integrity
}

// OrderItem represents the product data tied to a specific order.
// This data is fetched by JOINing 'order_product' (for quantity/price) and 'products' (for name).
type OrderItem struct {
    ProductID int     `json:"product_id"`
    Name      string  `json:"product_name"` // From the 'products' table
    Quantity  int     `json:"quantity"`     // From the 'order_product' junction table
    UnitPrice float64 `json:"unit_price"`   // From the 'order_product' junction table
}

type OrderDetails struct {
    OrderID      int            `json:"order_id"`
    OrderedDate  time.Time      `json:"ordered_date"`
    DeliverDate  *time.Time     `json:"delivered_date"`
    Status       string         `json:"status"`
    UserID       int            `json:"user_id"`
    
    // The list of products specific to this order
    Items        []OrderItem    `json:"items"` 
}
