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
