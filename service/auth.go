package service

import (
  "log"
  "time"
  "context"
  "errors"
  "os"

  "golang.org/x/crypto/bcrypt"
  "github.com/golang-jwt/jwt/v5"

  "codeberg.org/noreng-br/models"
)

var ErrNotFound = errors.New("User not found")
var ErrPassword = errors.New("Password does not match")

func (s *Service) Login(ctx context.Context, auth *models.Auth) (string, error) {
  log.Println(ctx, auth)
  // attempot to get user by its username
  usr, err := s.Repository.UserRepo.GetUserByUsername(ctx, auth.Login)
  if err != nil {
    // attempt to get user by its email
    usr, err = s.Repository.UserRepo.GetUserByEmail(ctx, auth.Login)
    if err != nil {
      return "", ErrNotFound
    }
  }
  log.Println("===============User FOUND+========================")
  log.Println(usr)
  log.Println("================================================+")

  // compare passwords
  checkPassword := func (password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil 
  }(auth.Password, usr.Password)
  if !checkPassword {
    log.Println("Passwords don't match")
    log.Println(usr.Password)
    return "", ErrPassword
  }

 log.Println("Passwords do match ==== successs")
 log.Println("User Id; ", usr.ID)

  // check whether or not the user is an admin, specific privileges are granted to admin users
  var role string
  if usr.IsAdmin {
    role = "admin" 
  } else {
    role = "user"
  }

  claims := &models.JWTCustomClaims{
        UserID: usr.ID,
        Role: role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // Token expires in 1 hour
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
  } 

  secretKey := os.Getenv("JWT_SECRET")

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  t, err := token.SignedString([]byte(secretKey))
  if err != nil {
      return "", err
  }

  return t, nil
}
