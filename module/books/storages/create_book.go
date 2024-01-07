package bookstorage

import (
	"context"

	bookmodel "local-app/module/books/models"
)

// CreateBook creates a new book entry in the database.
func (s *dbStorage) CreateBook(ctx context.Context, data *bookmodel.Book) error {
	// Attempt to create a new entry in the database using the provided Book data
	if err := s.db.Create(data).Error; err != nil {
		return err // Return the error if there's any issue in creating the book entry
	}
	return nil // Return nil indicating success if no error occurred
}
