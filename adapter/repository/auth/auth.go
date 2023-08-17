package auth

import (
	"awesomeProject/domain"
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase/repository/auth"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.AuthRepository {
	return &authRepository{db: db}
}

func (aR *authRepository) LoginUser(loginForm *domain.LoginFormDTO) (*domain.User, error) {
	var alreadyExistUser *domain.User

	var (
		loginRes *domain.User
		err      error
	)
	loginRes, err = nil, errors.New(fmt.Sprintf("Email or password is incorrect"))

	existUserErr := aR.db.Model(&domain.User{}).Where("email = ?", loginForm.Email).First(&alreadyExistUser).Error
	if existUserErr != nil || !utils.CheckPasswordHash(loginForm.Password, *alreadyExistUser.Password) {
		return loginRes, err
	}

	loginRes, err = alreadyExistUser, nil
	return loginRes, err
}

func (aR *authRepository) CreateUser(user *domain.User) (*domain.User, error) {
	var dbUser *domain.User
	if err := aR.db.Model(&domain.User{}).Where("email = ? OR phone = ?", user.Email, user.Phone).First(&dbUser).Error; err == nil {
		if dbUser.Email == user.Email {
			return nil, errors.New(fmt.Sprintf("A user with email %s already existing", user.Email))
		}
		if dbUser.Phone == user.Phone {
			return nil, errors.New(fmt.Sprintf("A user with phone %d already existing", user.Phone))
		}
	}

	if err := aR.db.Model(&domain.User{}).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
