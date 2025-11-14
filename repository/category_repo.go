package repository

import (
    "context"

    "codeberg.org/noreng-br/models"
)

type CategoryRepository interface {
  CreateCategory(ctx context.Context, name string) (models.Category, error)
  DeleteCategory(ctx context.Context, categoryId string) (error)
  ListCategories(ctx context.Context) ([]models.Category, error)
}
