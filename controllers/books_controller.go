package controllers

import (
	"ginLibrary/models"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var bookService BookService // This would be set up in main.go and injected here

// SetBookService sets the book service for the controllers
func SetBookService(service BookService) {
	bookService = service
}

// BookService is the interface for our book service
type BookService interface {
	AddBook(book models.Book) error
	DeleteBook(bookName string) error
	GetAllBooks(userType string) ([]models.Book, error)
}

// AddBook handles adding a new book
func AddBook(c *gin.Context) {
	// Extract user type from JWT stored in the context by middleware
	tokenInterface, exists := c.Get("userToken")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
		return
	}

	tokenClaims := tokenInterface.(*jwt.Token).Claims.(jwt.MapClaims)
	userType := tokenClaims["userType"].(string)

	if userType != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
		return
	}

	if book.BookName == "" || book.Author == "" || book.PublicationYear == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	err := bookService.AddBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book added", "book": book})
}

// DeleteBook handles deleting a book
func DeleteBook(c *gin.Context) {
	// Extract user type from JWT stored in the context by middleware
	tokenInterface, exists := c.Get("userToken")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
		return
	}

	tokenClaims := tokenInterface.(*jwt.Token).Claims.(jwt.MapClaims)
	userType := tokenClaims["userType"].(string)

	if userType != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	bookName := c.Query("bookName")
	if bookName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book name required"})
		return
	}

	err := bookService.DeleteBook(strings.TrimSpace(strings.ToLower(bookName)))
	if err != nil {
		if err.Error() == "no matching record found" {
			c.JSON(http.StatusOK, gin.H{"message": "No matching record found. No record deleted"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "1 record deleted"})
}

func GetAllBooks(c *gin.Context) {
	// Extract user type from JWT stored in the context by middleware
	tokenInterface, exists := c.Get("userToken")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
		return
	}

	tokenClaims, ok := tokenInterface.(*jwt.Token).Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error accessing token claims"})
		return
	}

	userType, ok := tokenClaims["userType"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User type is invalid or not found in token"})
		return
	}

	books, err := bookService.GetAllBooks(userType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

// curl -H "Authorization: Bearer {token}" http://localhost:8080/home
// curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer {token}" -d '{"bookName":"New Book", "author":"Author Name", "publicationYear":2021}' http://localhost:8080/addBook
// curl -X DELETE -H "Authorization: Bearer {token}" http://localhost:8080/deleteBook?bookName=New%20Book
