package controllers

import (
	"ginLibrary/models"
	"ginLibrary/services"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Login handles the login functionality
func Login(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	log.Printf("Attempting to authenticate user: %s", credentials.Username)
	user, authenticated := services.AuthenticateUser(credentials.Username, credentials.Password)
	if !authenticated {
		log.Printf("Authentication failed for user: %s", credentials.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	log.Printf("User authenticated, generating JWT for user: %s", user.Username)
	token, err := services.GenerateJWT(*user)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	expirationHours, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_HOURS"))
	if err != nil {
		log.Fatalf("Failed to convert TOKEN_EXPIRATION_HOURS to integer: %v", err)
	}
	expirationTime := time.Now().Add(time.Duration(expirationHours) * time.Hour)

	c.JSON(http.StatusOK, gin.H{"Bearer Token": token, "Expires At": expirationTime})
}
