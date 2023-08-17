package usecase

import (
	"awesomeProject/domain"
	"awesomeProject/usecase/repository"
	"awesomeProject/usecase/repository/auth"
	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	Login(loginForm *domain.LoginFormDTO) (*domain.LoginResponseDTO, error)
	Register(user *domain.User) (*domain.RegisterResponseDTO, error)
	Logout(refreshToken string) error
}

type authUseCase struct {
	authRepository   auth.AuthRepository
	tokensRepository repository.TokensRepository
	dbRepository     repository.DBRepository
}

func NewAuthUseCase(aR auth.AuthRepository, tR repository.TokensRepository, dbR repository.DBRepository) Auth {
	return &authUseCase{
		authRepository:   aR,
		tokensRepository: tR,
		dbRepository:     dbR,
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

func (auc *authUseCase) Login(loginForm *domain.LoginFormDTO) (*domain.LoginResponseDTO, error) {
	user, loginUserErr := auc.authRepository.LoginUser(loginForm)
	if loginUserErr != nil {
		return nil, loginUserErr
	}

	tokens, tokensErr := auc.tokensRepository.GenerateTokens(domain.NewUserDTO(user))
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
		User:         *domain.NewUserDTO(user),
	}, nil
}

func (auc authUseCase) Register(user *domain.User) (*domain.RegisterResponseDTO, error) {
	usr, createUserErr := auc.authRepository.CreateUser(user)
	if createUserErr != nil {
		return nil, createUserErr
	}

	tokens, tokensErr := auc.tokensRepository.GenerateTokens(domain.NewUserDTO(usr))

	if tokensErr != nil {
		return nil, tokensErr
	}

	_, err := auc.tokensRepository.SaveRefreshToken(usr.ID, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &domain.RegisterResponseDTO{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         *domain.NewUserDTO(usr),
	}, nil
}
