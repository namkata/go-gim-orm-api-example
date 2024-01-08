package bookstorage

import (
	"context"
	"local-app/common"
	bookmodel "local-app/modules/books/models"
)

// ListBook retrieves a list of books based on provided conditions and pagination settings.
// It takes a context, conditions for filtering, and pagination details.
func (s *dbStorage) ListBook(
	ctx context.Context,
	condition map[string]interface{},
	filter *bookmodel.Filter,
	paging *common.Paging,
) ([]bookmodel.Book, error) {
	// Calculate the offset based on the page number and limit
	offset := (paging.Page - 1) * paging.Limit

	var result []bookmodel.Book
	query := s.db.Table(bookmodel.Book{}.TableName()).Where(condition)

	// Formulate the query based on the provided filter criteria
	if v := filter; v != nil {
		if v.Name != "" {
			query = query.Where("name LIKE ?", "%"+v.Name+"%")
		}
		// Adding price range filtering logic
		if filter.Price.Min != 0 || filter.Price.Max != 0 {
			query = query.Where("price BETWEEN ? AND ?", filter.Price.Min, filter.Price.Max)
		}
	}

	// Count total records
	if err := query.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	// Retrieve records with pagination and ordering
	if err := query.
		Limit(paging.Limit).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
