package main

import (
	"GOFILEGO/configs"
	"GOFILEGO/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configure the Gin mode based on environment variable
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.ReleaseMode // default to release mode if not set
	}
	gin.SetMode(mode)

	router := SetupAppRouter()

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
func SetupAppRouter() *gin.Engine {
	db := configs.Connection()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	api := router.Group("api/v1")

	file := api.Group("/file")

	routes.InitAuthRoutes(db, api)

	routes.InitFileRoutes(db, file)Í
	return router
}
