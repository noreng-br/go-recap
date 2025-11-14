package repository

import (
    "context"

    "codeberg.org/noreng-br/models"
    )

type ProductRepository interface {
  CreateProduct(ctx context.Context, product models.Product) (models.Product, error)
  GetProducts(ctx context.Context) ([]models.ProductWithCategories, error)
  GetProductById(ctx context.Context, productId string) (models.ProductWithCategories, error)
  AddCategoriesToProduct(ctx context.Context, productID int, categoryIDs []int) error
  GetCategoriesByProductID(ctx context.Context, productID int) ([]models.Category, error)
  UpdateProduct(ctx context.Context, productId string, product models.Product) (models.Product, error)
  DeleteProduct(ctx context.Context, productId string) ( error)
}
