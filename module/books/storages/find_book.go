package bookstorage

import (
	"context"
	bookmodel "local-app/module/books/models"

	"gorm.io/gorm"
)

// FindBook finds a book based on certain conditions in the database.
func (s *dbStorage) FindBook(ctx context.Context, condition map[string]interface{}) (*bookmodel.Book, error) {
	var bookData bookmodel.Book

	// Query the database to find a book matching the given conditions
	if err := s.db.Where(condition).First(&bookData).Error; err != nil {
		// If no matching record is found, return a custom "Book not found" error
		if err == gorm.ErrRecordNotFound {
			return nil, bookmodel.ErrBookNotFound
		}

		// Return other errors encountered during the query
		return nil, err
	}

	// Return the found book data
	return &bookData, nil
}
