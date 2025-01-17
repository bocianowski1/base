package services

type IAuthService interface {
	Login(email, password string) (string, error)
	// Logout(token string) error
}

type AuthService struct {
	// repo repo.IUserRepo
	// logger repo.ILogger
	// cache  repo.ICache
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(email, password string) (string, error) {
	return "valid token", nil
}
