package usecase

import (
	"awesomeProject/domain"
	"awesomeProject/usecase/repository"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type TokensAuth interface {
	RefreshToken(refreshToken string) (*domain.RefreshTokenResponseDTO, error)
}

type tokensUseCase struct {
	tokensRepository repository.TokensRepository
	dbRepository     repository.DBRepository
}

func NewTokensUseCase(tR repository.TokensRepository, dbR repository.DBRepository) TokensAuth {
	return &tokensUseCase{
		tokensRepository: tR,
		dbRepository:     dbR,
	}
}

func (auc *tokensUseCase) RefreshToken(refreshToken string) (*domain.RefreshTokenResponseDTO, error) {
	if len(refreshToken) == 0 {
		return nil, jwt.ErrTokenUnverifiable
	}

	user, validateTokenErr := auc.tokensRepository.ValidateRefreshToken(refreshToken)

	if validateTokenErr != nil {
		return nil, errors.New("unauthorized error")
	}

	_, tokenErrExist := auc.tokensRepository.FindRefreshToken(user.ID)
	if tokenErrExist != nil {
		return nil, errors.New("unauthorized error")
	}

	userDTO := domain.NewUserDTO(user)
	tokens, tokensErr := auc.tokensRepository.GenerateTokens(userDTO)
	if tokensErr != nil {
		return nil, tokensErr
	}
	_, err := auc.tokensRepository.SaveRefreshToken(user.ID, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &domain.RefreshTokenResponseDTO{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
