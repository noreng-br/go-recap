package repository

import (
    "context"
    "codeberg.org/noreng-br/models"
)

type OrderRepository interface {
  CreateOrder(ctx context.Context, req models.OrderRequest) (int, error)
  GetUserOrders(ctx context.Context, userID int) ([]models.OrderDetails, error)
  ListOrders(ctx context.Context) ([]models.OrderDetails, error)
}
