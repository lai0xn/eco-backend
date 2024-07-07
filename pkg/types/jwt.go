package types

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
  Name string
  Email string
  jwt.RegisteredClaims
}
