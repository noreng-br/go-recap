package models

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
  UserID string `json:"user_id"`
  Role string `json:"role"`
  jwt.RegisteredClaims
}
