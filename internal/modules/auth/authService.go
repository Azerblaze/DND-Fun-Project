package auth

// Declare AuthService interface
type IAuthService interface {
	Test(text string) string
}

// Declare AuthService type
type AuthService struct {
	authRepository IAuthRepository
}

// Function to initialize AuthService new instace
func NewAuthService(authRepository IAuthRepository) IAuthService {
	return &AuthService{authRepository}
}

func (authService *AuthService) Test(text string) string {
	result := authService.authRepository.Test(text)

	return result
}
