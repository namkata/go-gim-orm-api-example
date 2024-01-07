package bookstorage

import (
	"context"

	bookmodel "local-app/module/books/models"
)

// ListBook retrieves a list of books based on provided conditions and pagination settings.
// It takes a context, conditions for filtering, and pagination details.
func (s *dbStorage) ListBook(
	ctx context.Context,
	condition map[string]interface{},
	paging *bookmodel.DataPaging,
) ([]bookmodel.Book, error) {
	// Calculate the offset based on the page number and limit
	offset := (paging.Page - 1) * paging.Limit

	var result []bookmodel.Book

	// Count the total number of records that match the given conditions
	// Perform the query, applying conditions, pagination, and ordering by ID in descending order
	if err := s.db.
		Table(bookmodel.Book{}.TableName()).
		Where(condition).
		Count(&paging.Total).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
