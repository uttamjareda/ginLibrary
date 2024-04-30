package repositories

import (
	"encoding/csv"
	"errors"
	"ginLibrary/models"
	"os"
	"strconv"
	"strings"
)

type BookRepository struct {
	regularFilePath string
	adminFilePath   string
}

func NewBookRepository(regularPath, adminPath string) *BookRepository {
	return &BookRepository{
		regularFilePath: regularPath,
		adminFilePath:   adminPath,
	}
}

// GetAllBooks reads all books from both regular and admin CSV files
func (repo *BookRepository) GetAllBooks() ([]models.Book, error) {
	books, err := repo.readBooks(repo.regularFilePath)
	if err != nil {
		return nil, err
	}
	adminBooks, err := repo.readBooks(repo.adminFilePath)
	if err != nil {
		return nil, err
	}
	return append(books, adminBooks...), nil
}

// GetRegularBooks reads books from the regular user CSV file
func (repo *BookRepository) GetRegularBooks() ([]models.Book, error) {
	return repo.readBooks(repo.regularFilePath)
}

// AddBook adds a new book to the regular user CSV file
func (repo *BookRepository) AddBook(book models.Book) error {
	// file, err := os.OpenFile(repo.regularFilePath, os.O_WRONLY|os.O_APPEND, 0644)
	file, err := os.OpenFile(repo.regularFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{book.BookName, book.Author, strconv.Itoa(book.PublicationYear)}
	if err := writer.Write(record); err != nil {
		return err
	}
	return nil
}

// DeleteBook deletes a book from the regular user CSV file based on book name
func (repo *BookRepository) DeleteBook(bookName string) error {
	tempFilePath := "assets/temp.csv"
	originalFile, err := os.Open(repo.regularFilePath)
	if err != nil {
		return err
	}
	defer originalFile.Close()

	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	csvReader := csv.NewReader(originalFile)
	csvWriter := csv.NewWriter(tempFile)
	defer csvWriter.Flush()

	found := false
	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}
		if strings.ToLower(record[0]) == strings.ToLower(bookName) {
			found = true
			continue
		}
		if err := csvWriter.Write(record); err != nil {
			return err
		}
	}

	if !found {
		return errors.New("no matching record found")
	}

	// Replace old file with the updated one
	os.Remove(repo.regularFilePath)
	os.Rename(tempFilePath, repo.regularFilePath)

	return nil
}

func (repo *BookRepository) readBooks(filePath string) ([]models.Book, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var books []models.Book
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records[1:] { // Skip header
		year, _ := strconv.Atoi(record[2])
		books = append(books, models.Book{BookName: record[0], Author: record[1], PublicationYear: year})
	}

	return books, nil
}
