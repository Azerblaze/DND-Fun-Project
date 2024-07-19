package routes

import (
	"dnd-fun-be/internal/db"
	"dnd-fun-be/internal/modules/auth"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {
	// Group api using /api url
	api := app.Group("/api")

	DB := db.InitDatabase()

	// Add Route based on modules
	auth.AuthRoutes(api, DB)
}
