package usecase

import (
	"awesomeProject/domain"
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase/repository"
	"awesomeProject/usecase/repository/auth"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Auth interface {
	Login(loginForm *domain.LoginFormDTO) (*domain.LoginResponseDTO, error)
	Register(user *domain.User) (*domain.RegisterResponseDTO, error)
	Logout(refreshToken string) error
	Activate(activationLink string) (*domain.LoginResponseDTO, error)
}

type authUseCase struct {
	authRepository   auth.AuthRepository
	tokensRepository repository.TokensRepository
	dbRepository     repository.DBRepository
	mailService      utils.MailService
}

func NewAuthUseCase(aR auth.AuthRepository, tR repository.TokensRepository, dbR repository.DBRepository, ms utils.MailService) Auth {
	return &authUseCase{
		authRepository:   aR,
		tokensRepository: tR,
		dbRepository:     dbR,
		mailService:      ms,
	}
}

func (auc *authUseCase) Logout(refreshToken string) error {
	user, validateErr := auc.tokensRepository.ValidateRefreshToken(refreshToken)
	if validateErr != nil {
		return jwt.ErrTokenExpired
	}
	err := auc.tokensRepository.RemoveRefreshToken(user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (auc *authUseCase) Activate(activationLink string) (*domain.LoginResponseDTO, error) {
	user, activateErr := auc.authRepository.ActivateUser(activationLink)
	if activateErr != nil {
		return nil, activateErr
	}

	tokens, tokensErr := auc.tokensRepository.GenerateTokens(user)
	if tokensErr != nil {
		return nil, tokensErr
	}

	_, err := auc.tokensRepository.SaveRefreshToken(user.ID, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponseDTO{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         *user,
	}, nil
}

func (auc *authUseCase) Login(loginForm *domain.LoginFormDTO) (*domain.LoginResponseDTO, error) {
	user, loginUserErr := auc.authRepository.LoginUser(loginForm)
	if loginUserErr != nil {
		return nil, loginUserErr
	}

	tokens, tokensErr := auc.tokensRepository.GenerateTokens(user)
	if tokensErr != nil {
		return nil, tokensErr
	}

	_, err := auc.tokensRepository.SaveRefreshToken(user.ID, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponseDTO{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         *user,
	}, nil
}

func (auc authUseCase) Register(user *domain.User) (*domain.RegisterResponseDTO, error) {
	hashPass, _ := utils.HashPassword(*user.Password)
	user.Password = &hashPass

	activationLink := uuid.Must(uuid.NewRandom()).String()

	usr, createUserErr := auc.authRepository.CreateUser(user, activationLink)
	if createUserErr != nil {
		return nil, createUserErr
	}

	tokens, tokensErr := auc.tokensRepository.GenerateTokens(usr)
	if tokensErr != nil {
		return nil, tokensErr
	}

	fmt.Println(usr.ID)
	_, err := auc.tokensRepository.SaveRefreshToken(usr.ID, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}

	go auc.mailService.SendMail(fmt.Sprintf(
		"SUCCESSFULY REGISTRATION! PLEASE GO http://localhost:3000/?server-action=user-activation&value=%s TO ACTIVATE YOUR ACCOUNT",
		activationLink),
	)

	return &domain.RegisterResponseDTO{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         *usr,
	}, nil
}
