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
	router := gin.Default()

	// Connect to the database
	db, err := configs.Connection()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	api := router.Group("/api/v1")
	routes.InitAuthRoutes(db, api)

	return router
}
