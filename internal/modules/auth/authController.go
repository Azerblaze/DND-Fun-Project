package auth

import (
	"dnd-fun-be/internal/modules/auth/dto"

	"github.com/gin-gonic/gin"
)

// Declare AuthController Type
type AuthController struct {
	authService IAuthService
}

// Function to initialize new instance of controller
func NewAuthController(authService IAuthService) *AuthController {
	return &AuthController{authService}
}

func (authController *AuthController) Test(ctx *gin.Context) {
	var text = &dto.TestText{}

	// Bind Context to variable
	err := ctx.Bind(text)
	if err != nil {
		panic(err)
	}

	// Call the Service
	result := authController.authService.Test(text.Text)

	ctx.JSON(200, gin.H{
		"text": result,
	})
}
