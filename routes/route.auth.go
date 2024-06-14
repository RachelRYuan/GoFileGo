package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	loginController "GOFILEGO/controllers/auth-controllers/login"
	registerController "GOFILEGO/controllers/auth-controllers/register"
	loginHandler "GOFILEGO/handlers/auth-handlers/login"
	registerHandler "GOFILEGO/handlers/auth-handlers/register"
)

func InitAuthRoutes(db *gorm.DB, route *gin.RouterGroup) {
	// Initialize login components
	loginRepository := loginController.NewRepositoryLogin(db)
	loginService := loginController.NewServiceLogin(loginRepository)
	loginHandler := loginHandler.NewHandlerLogin(loginService)

	// Initialize register components
	registerRepository := registerController.NewRegisterRepository(db)
	registerService := registerController.NewRegisterService(registerRepository)
	registerHandler := registerHandler.NewHandlerRegister(registerService)

	// Define routes
	route.POST("/login", loginHandler.LoginHandler)
	route.POST("/register", registerHandler.RegisterHandler)
}
