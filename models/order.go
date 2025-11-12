package models

import (
  "time"
)

type Order struct {
  ID  string  `json:"id"`
  OrderedDate time.Time  `json:"ordered_date"`
  DeliverDate time.Time  `json:"deliver_date"`
  Products []Product `json:"products"`
}
