package domain

type LoginFormDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"-"`
	User         UserDTO `json:"user"`
}

type RegisterResponseDTO struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"-"`
	User         UserDTO `json:"user"`
}

type RefreshTokenResponseDTO struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"-"`
	User         UserDTO `json:"user"`
}
