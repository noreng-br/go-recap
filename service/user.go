package service

import (
  "context"
  "codeberg.org/noreng-br/models"
)

type UserService interface {
  CreateUser(user models.User) (models.User, error)
}

func (s *Service) CreateUser(ctx context.Context, user models.User) (models.User, error) {
  user, err := s.Repository.UserRepo.CreateUser(ctx, user)
  if err != nil {
    return models.User{}, err
  }
  
  return user, nil
}
