package repository

import (
  "context"
  "codeberg.org/noreng-br/models"
)

type UserRepository interface {
  CreateUser(ctx context.Context, user models.User) (models.User, error)
  GetUserByUsername(ctx context.Context, username string) (models.User, error)
  GetUserByEmail(ctx context.Context, email string) (models.User, error)
  GetUsers(ctx context.Context) ([]models.User, error)
}
