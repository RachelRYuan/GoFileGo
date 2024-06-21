package routes

import (
	"GOFILEGO/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	filecontrollers "GOFILEGO/controllers/file-controllers"
	filehandlers "GOFILEGO/handlers/file-handlers"
)

// InitFileRoutes initializes the routes for file operations.
func InitFileRoutes(db *gorm.DB, route *gin.RouterGroup) {
	fileRepository := filecontrollers.NewFileRepository(db)
	fileService := filecontrollers.NewFileService(fileRepository)
	fileHandlers := filehandlers.NewCreateHandler(fileService)

	// Add auth middleware
	route.Use(middlewares.Auth())

	route.POST("/create", fileHandlers.CreateHandler)
	route.GET("/", fileHandlers.GetAllFilesHandler)
	route.DELETE("/:fileId", fileHandlers.DeleteHandler)
}
