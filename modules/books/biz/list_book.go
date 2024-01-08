package bookbiz

import (
	"context"
	"local-app/common"
	bookmodel "local-app/modules/books/models"
)

// ListBookStorage defines the interface for listing books.
type ListBookStorage interface {
	ListBook(
		ctx context.Context,
		condition map[string]interface{},
		filter *bookmodel.Filter,
		paging *common.Paging,
	) ([]bookmodel.Book, error)
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
func (biz *listBookBiz) ListBook(
	ctx context.Context,
	filter *bookmodel.Filter,
	paging *common.Paging,
) ([]bookmodel.Book, error) {
	// Delegate the call to the storage layer
	result, err := biz.store.ListBook(ctx, map[string]interface{}{}, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
