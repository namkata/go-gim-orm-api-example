package bookstorage

import (
	"context"

	bookmodel "local-app/module/books/models"
)

// UpdateBook updates the book information based on a specific condition with new data.
func (s *dbStorage) UpdateBook(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *bookmodel.Book,
) error {
	// Updates the book information in the database based on the provided condition and data
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
