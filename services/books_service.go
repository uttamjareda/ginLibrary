package services

import (
	"errors"
	"ginLibrary/models"
	"ginLibrary/repositories"
	"log"
	"strings"
)

type BookService struct {
	bookRepo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
	return &BookService{
		bookRepo: repo,
	}
}

// GetAllBooks returns all books for a given user type
func (s *BookService) GetAllBooks(userType string) ([]models.Book, error) {
	var books []models.Book
	var err error

	log.Printf("userType: %v", userType)
	if userType == "admin" {
		log.Printf("userType admin: %v", userType)
		books, err = s.bookRepo.GetAllBooks()
	} else {
		log.Printf("userType regular: %v", userType)
		books, err = s.bookRepo.GetRegularBooks()
	}

	if err != nil {
		return nil, err
	}

	log.Printf("books: %v", books)

	return books, nil
}

// AddBook adds a new book to the system
func (s *BookService) AddBook(book models.Book) error {
	// Validate the book data (this could be more extensive)
	if book.BookName == "" || book.Author == "" || book.PublicationYear == 0 {
		return errors.New("invalid book data")
	}

	return s.bookRepo.AddBook(book)
}

// DeleteBook removes a book from the system
func (s *BookService) DeleteBook(bookName string) error {
	if bookName == "" {
		return errors.New("book name must be provided")
	}

	bookName = strings.TrimSpace(strings.ToLower(bookName))
	return s.bookRepo.DeleteBook(bookName)
}
