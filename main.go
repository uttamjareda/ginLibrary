package main

import (
	"ginLibrary/controllers"
	"ginLibrary/repositories"
	"ginLibrary/routes"
	"ginLibrary/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Set up file paths from .env or default
	adminFilePath := os.Getenv("ADMIN_CSV_PATH")
	if adminFilePath == "" {
		adminFilePath = "assets/adminUser.csv"
	}
	regularFilePath := os.Getenv("REGULAR_CSV_PATH")
	if regularFilePath == "" {
		regularFilePath = "assets/regularUser.csv"
	}

	// Initialize the repository
	bookRepo := repositories.NewBookRepository(regularFilePath, adminFilePath)
	bookService := services.NewBookService(bookRepo)

	// Set up the router
	router := gin.Default()

	// Inject the bookService into the controllers
	controllers.SetBookService(bookService) // This method needs to be implemented in controllers package

	// Initialize routes
	routes.InitializeRoutes(router)

	// Start the server
	router.Run(":8080") // Run on http://localhost:8080
}
