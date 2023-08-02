package repository

import (
	"awesomeProject/domain"
	"awesomeProject/usecase/repository"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"os"
	"time"
)

type tokensRepository struct {
	db *gorm.DB
}

func NewTokensRepository(db *gorm.DB) repository.TokensRepository {
	return &tokensRepository{db: db}
}

func (tR tokensRepository) FindRefreshToken(refreshToken string) (*domain.Token, error) {
	var tokenData *domain.Token
	if err := tR.db.Model(&domain.Token{}).Where("refresh_token = ?", refreshToken).First(&tokenData).Error; err != nil {
		return nil, err
	}

	return tokenData, nil
}

func (tR tokensRepository) RemoveRefreshToken(refreshToken string) (*domain.Token, error) {
	err := tR.db.Model(&domain.Token{}).Where("refresh_token = ?", refreshToken).Delete(&domain.Token{}).Error

	return nil, err
}

func validateToken(tokenString string, tokenKey []byte) (*domain.UserDTO, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*domain.JWTPayload); ok && token.Valid {
		return &claims.User, nil
	} else {
		return nil, jwt.ErrTokenExpired
	}
}

func ValidateAccessToken(accessToken string) (*domain.UserDTO, error) {
	accessSecret := os.Getenv("SECRET_ACCESS_TOKEN")
	myAccessTokenKey := []byte(accessSecret)

	if user, err := validateToken(accessToken, myAccessTokenKey); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (tR tokensRepository) ValidateRefreshToken(refreshToken string) (*domain.User, error) {
	refreshSecret := os.Getenv("SECRET_REFRESH_TOKEN")
	myRefreshTokenKey := []byte(refreshSecret)
	if user, err := validateToken(refreshToken, myRefreshTokenKey); err != nil {
		return nil, err
	} else {
		var userData *domain.User
		errUser := tR.db.Model(&domain.User{}).Where("id = ?", user.ID).First(&userData).Error
		if errUser != nil {
			return nil, errUser
		}
		return userData, nil
	}
}

func (tR tokensRepository) SaveRefreshToken(userID int, refreshToken string) (*domain.Token, error) {
	var tokenData *domain.Token

	if err := tR.db.Model(&domain.Token{}).Where("user_id = ?", userID).First(&tokenData).Update("refresh_token", refreshToken).Error; err == nil {
		return tokenData, nil
	}

	newTokenData := domain.Token{
		RefreshToken: refreshToken,
		UserId:       userID,
	}
	if err := tR.db.Model(&domain.Token{}).Create(&newTokenData).Error; err != nil {
		return nil, err
	}

	return &newTokenData, nil
}

func (tR tokensRepository) GenerateTokens(user *domain.UserDTO) (*domain.Tokens, error) {
	accessSecret := os.Getenv("SECRET_ACCESS_TOKEN")
	refreshSecret := os.Getenv("SECRET_REFRESH_TOKEN")

	myAccessTokenKey := []byte(accessSecret)
	myRefreshTokenKey := []byte(refreshSecret)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(10 * time.Minute),
		"user": user,
		"iat":  jwt.NewNumericDate(time.Now()),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(30 * time.Minute),
		"user": user,
	})

	accessTokenString, err := accessToken.SignedString(myAccessTokenKey)
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := refreshToken.SignedString(myRefreshTokenKey)
	if err != nil {
		return nil, err
	}

	return &domain.Tokens{
		RefreshToken: refreshTokenString,
		AccessToken:  accessTokenString,
	}, nil
}
