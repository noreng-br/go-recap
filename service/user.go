package service

import (
  "context"
  "log"

  "golang.org/x/crypto/bcrypt"

  "codeberg.org/noreng-br/models"
)

type UserService interface {
  CreateUser(ctx context.Context, user models.User) (models.User, error)
  GetUserByUsername(ctx context.Context, username string) (models.User, error)
  GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
  user, err := s.Repository.UserRepo.GetUserByUsername(ctx, username)
  if err != nil {
    log.Println("===============================")
    log.Println(err.Error())
    log.Println("===============================")
    return models.User{}, err
  }


  log.Println("======================================")
  log.Println("GET USERNAME SUCCESSFULL")
  log.Println("======================================")
  return user, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
  user, err := s.Repository.UserRepo.GetUserByEmail(ctx, email)
  if err != nil {
    log.Println("===============================")
    log.Println(err.Error())
    log.Println("===============================")
    return models.User{}, err
  }


  log.Println("======================================")
  log.Println("GET EMAIL SUCCESSFULL")
  log.Println("======================================")
  return user, nil
}


func (s *Service) CreateUser(ctx context.Context, user models.User) (models.User, error) {
  //need to hash the password
  makeHashPassword, err := func (password string) (string, error) {
    // The second argument is the 'cost' factor (the number of rounds).
    // bcrypt.DefaultCost is currently 10, which is generally sufficient.
    // Higher cost is slower but more secure.
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
      return "", err
    }
    // The returned bytes are the full hash string (cost + salt + hash)
    return string(bytes), nil
  }(user.Password)
  if err != nil {
    return models.User{}, err
  }

  user.Password = makeHashPassword
  
  user, err = s.Repository.UserRepo.CreateUser(ctx, user)
  if err != nil {
    log.Println("=================================================")
    log.Println("An error ocurred in the service.. could not insert user")
    log.Println(err.Error())
    log.Println("=================================================")
    return models.User{}, err
  }
  
  return user, nil
}
