package services

import (
	"fmt"
	"ginLibrary/db"
	"ginLibrary/models"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a JWT token for authenticated users
func GenerateJWT(user models.User) (string, error) {
	expirationHours, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_HOURS"))
	if err != nil {
		log.Fatalf("Failed to convert TOKEN_EXPIRATION_HOURS to integer: %v", err)
	}
	expirationTime := time.Now().Add(time.Duration(expirationHours) * time.Hour)

	claims := jwt.MapClaims{
		"userName":  user.Username,
		"userType":  user.UserType,
		"expiresAt": expirationTime.Unix(),
	}

	log.Printf("Claims set for JWT: %v", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Printf("Error signing JWT: %v", err)
		return "", err
	}

	return tokenString, nil
}

// AuthenticateUser checks if the username and password are correct
func AuthenticateUser(username, password string) (*models.User, bool) {
	for _, user := range db.Users {
		if user.Username == username && user.Password == password {
			log.Printf("User authenticated: %s", username)
			return &user, true
		}
	}
	log.Printf("User not found or password incorrect: %s", username)
	return nil, false
}

// VerifyTokenFromHeader extracts the JWT from the Authorization header and verifies it
func VerifyTokenFromHeader(authHeader string) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if tokenString == "" {
		return nil, fmt.Errorf("no token found")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Retrieve the JWT key from a secure location, such as environment variables
		jwtKey := []byte("YourJWTSecretHere") // Replace with actual secure retrieval of your key
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
