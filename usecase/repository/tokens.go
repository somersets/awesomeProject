package repository

import "awesomeProject/domain"

type TokensRepository interface {
	FindRefreshToken(refreshToken string) (*domain.Token, error)
	GenerateTokens(user *domain.UserDTO) (*domain.Tokens, error)
	SaveRefreshToken(userID int, refreshToken string) (*domain.Token, error)
	RemoveRefreshToken(userId int) error
	ValidateRefreshToken(refreshToken string) (*domain.User, error)
}
