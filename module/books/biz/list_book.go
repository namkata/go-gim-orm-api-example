package bookbiz

import (
	"context"
	bookmodel "local-app/module/books/models"
)

// ListBookStorage defines the interface for listing books.
type ListBookStorage interface {
	ListBook(ctx context.Context, condition map[string]interface{}, paging *bookmodel.DataPaging) ([]bookmodel.Book, error)
}

// listBookBiz represents the business logic for listing books.
type listBookBiz struct {
	store ListBookStorage
}

// NewListBookBiz creates a new instance of listBookBiz.
func NewListBookBiz(store ListBookStorage) *listBookBiz {
	return &listBookBiz{store: store}
}

// ListBook retrieves a list of books based on conditions and pagination.
func (biz *listBookBiz) ListBook(ctx context.Context, condition map[string]interface{}, paging *bookmodel.DataPaging) ([]bookmodel.Book, error) {
	// Delegate the call to the storage layer
	result, err := biz.store.ListBook(ctx, condition, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
