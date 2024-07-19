package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.RouterGroup, DB *gorm.DB) {

	// Initialize Controller, Service and Repository
	authRepository := NewAuthRepository(DB)
	authService := NewAuthService(authRepository)
	authHandler := NewAuthController(authService)
	authRouter := router.Group("/auth")

	// Start Routing
	authRouter.POST("/test", authHandler.Test)
}
