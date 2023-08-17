package domain

import "github.com/golang-jwt/jwt/v5"

type Tokens struct {
	RefreshToken string
	AccessToken  string
}

type JWTPayload struct {
	UserId int
	jwt.RegisteredClaims
}
