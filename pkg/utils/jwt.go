package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lai0xn/squid-tech/config"
	"github.com/lai0xn/squid-tech/pkg/types"
)

func GenerateJWT(id string, email string, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &types.Claims{
		ID:    id,
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	})
	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	claims := new(types.Claims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
