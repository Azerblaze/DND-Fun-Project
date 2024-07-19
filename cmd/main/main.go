package main

import (
	"dnd-fun-be/config"
	"dnd-fun-be/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	config.InitConfig()

	routes.Routes(app)

	app.Run(fmt.Sprintf(":%s", config.Cfg.APIPort))
}
