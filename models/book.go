// models/book.go
package models

// Book represents the structure of a book
type Book struct {
	BookName        string `json:"bookName"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationYear"`
}
