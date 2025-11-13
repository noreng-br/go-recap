package repository

import (
    "context"

    "codeberg.org/noreng-br/models"
)

type CategoryRepository interface {
  CreateCategory(ctx context.Context, name string) (models.Category, error)
}
