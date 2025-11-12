package repository

import (
  "context"
  "codeberg.org/noreng-br/models"
)

type UserRepository interface {
  CreateUser(ctx context.Context, user models.User) (models.User, error)
}
