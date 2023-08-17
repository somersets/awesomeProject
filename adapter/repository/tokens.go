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

func (tR tokensRepository) RemoveRefreshToken(userId int) error {
	var user *domain.User
	if err := tR.db.Model(&domain.User{}).Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}
	err := tR.db.Model(&domain.Token{}).Where("user_id = ?", userId).Delete(&domain.Token{}).Error
	if err != nil {
		return err
	}
	return nil
}

func validateToken(tokenString string, tokenKey []byte) (*int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*domain.JWTPayload); ok && token.Valid {
		return &claims.UserId, nil
	} else {
		return nil, jwt.ErrTokenExpired
	}
}

func ValidateAccessToken(accessToken string) (*int, error) {
	accessSecret := os.Getenv("SECRET_ACCESS_TOKEN")
	myAccessTokenKey := []byte(accessSecret)

	if id, err := validateToken(accessToken, myAccessTokenKey); err != nil {
		return nil, err
	} else {
		return id, nil
	}
}

func (tR tokensRepository) ValidateRefreshToken(refreshToken string) (*domain.User, error) {
	refreshSecret := os.Getenv("SECRET_REFRESH_TOKEN")
	myRefreshTokenKey := []byte(refreshSecret)
	if id, err := validateToken(refreshToken, myRefreshTokenKey); err != nil {
		return nil, err
	} else {
		var userData *domain.User
		errUser := tR.db.Model(&domain.User{}).Where("id = ?", id).First(&userData).Error
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

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.JWTPayload{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.JWTPayload{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
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
