package auth

import "gorm.io/gorm"

// Declare AuthRepository Type
type AuthRepository struct {
	db *gorm.DB
}

// Declare AuthRepository interface
type IAuthRepository interface {
	Test(text string) string
}

// Function to initialize AuthRepository instance
func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db}
}

func (authRepository *AuthRepository) Test(text string) string {
	return text
}
