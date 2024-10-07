package router

import (
	"database/sql"
	"go-login-crud/internal/controller"
	"go-login-crud/internal/middleware"
	"go-login-crud/internal/repository"
	"go-login-crud/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateUserRouter(r *gin.Engine, db *sql.DB) {

	// Repository
	userRepo := repository.CreateUserRepository(db)

	//Service
	userService := service.CreateUserService(userRepo)
	authService := service.CreateAuthService(userRepo)
	
	// Controller
	userController := controller.CreateUserController(userService)
	authController := controller.CreateAuthController(authService)

	auth := r.Group("/api/v1/auth") 
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", userController.UserRegister)
	}

	jwt := middleware.Authenticate()
	user := r.Group("/api/v1/user").Use(jwt)
	{
		user.POST("/", userController.GetAllData)
		user.POST("/update", userController.UserUpdate)
		user.POST("/delete", userController.UserDelete)
	}
}