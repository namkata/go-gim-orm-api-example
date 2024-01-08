package bookstorage

import (
	"context"

	bookmodel "local-app/modules/books/models"
)

// DeleteBook deletes book data from the database based on specified conditions.
func (s *dbStorage) DeleteBook(
	ctx context.Context,
	condition map[string]interface{},
) error {
	// Construct and execute the deletion query based on the conditions
	if err := s.db.Table(bookmodel.Book{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
