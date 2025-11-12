package service

import (
  "codeberg.org/noreng-br/repository"
  )

type Service struct {
  Repository repository.Repositories
}

func NewService(r repository.Repositories) (*Service, error) {
  return &Service{
    Repository: r,
  }, nil
}
