package service

import (
  "context"
  "log"
  "codeberg.org/noreng-br/models"
)

type UserService interface {
  CreateUser(user models.User) (models.User, error)
}

func (s *Service) CreateUser(ctx context.Context, user models.User) (models.User, error) {
  user, err := s.Repository.UserRepo.CreateUser(ctx, user)
  if err != nil {
    log.Println("=================================================")
    log.Println("An error ocurred in the service.. could not insert user")
    log.Println(err.Error())
    log.Println("=================================================")
    return models.User{}, err
  }
  
  return user, nil
}
