package repository

import (
    "context"

    "codeberg.org/noreng-br/models"
    )

type ProductRepository interface {
  CreateProduct(ctx context.Context, product models.Product) (models.Product, error)
  GetProducts(ctx context.Context) ([]models.Product, error)
  GetProductById(ctx context.Context, productId string) (models.Product, error)
}
