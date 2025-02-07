package main

import (
	"github.com/gin-gonic/gin"
	"golang-web-server/config"
	"golang-web-server/routes"
	"os"
)

func main() {
	// setup config
	config.SetupConfig()

	// setup routes
	r := gin.Default()
	routes.SetupUserRoutes(r)
	routes.SetupProductRoutes(r)

	// start server
	r.Run(":" + os.Getenv("APP_PORT"))
}
