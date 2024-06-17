package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	filecontrollers "GOFILEGO/controllers/file-controllers"
	filehandlers "GOFILEGO/handlers/file-handlers"
)

// InitFileRoutes initializes the routes for file operations.
func InitFileRoutes(db *gorm.DB, route *gin.RouterGroup) {
	fileRepository := filecontrollers.NewFileRepository(db)
	fileService := filecontrollers.NewFileService(fileRepository)
	fileHandlers := filehandlers.NewHandler(fileService)

	route.POST("/create", fileHandlers.CreateFileHandler)
	route.GET("/", fileHandlers.GetAllFilesHandler)
	route.DELETE("/:fileId", fileHandlers.DeleteFileHandler)
}
