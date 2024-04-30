package routes

import (
	"ginLibrary/controllers"
	"ginLibrary/middleware"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes sets up the routes for the application
func InitializeRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.GET("/home", middleware.JWTAuthMiddleware(), controllers.GetAllBooks)
	router.POST("/addBook", middleware.JWTAuthMiddleware(), controllers.AddBook)
	router.DELETE("/deleteBook", middleware.JWTAuthMiddleware(), controllers.DeleteBook)
}
