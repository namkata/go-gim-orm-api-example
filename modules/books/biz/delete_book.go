package bookbiz

import (
	"context"

	bookmodel "local-app/modules/books/models"
)

// DeleteBookStorage specifies the methods needed for deleting book data.
type DeleteBookStorage interface {
	FindBook(ctx context.Context, condition map[string]interface{}) (*bookmodel.Book, error)
	DeleteBook(ctx context.Context, condition map[string]interface{}) error
}

// deleteBookBiz handles the business logic for deleting book data.
type deleteBookBiz struct {
	store DeleteBookStorage
}

// NewDeleteBookBiz creates a new deleteBookBiz instance.
func NewDeleteBookBiz(store DeleteBookStorage) *deleteBookBiz {
	return &deleteBookBiz{store: store}
}

// DeleteBook deletes book data using the provided condition from the storage layer.
func (biz *deleteBookBiz) DeleteBook(ctx context.Context, condition map[string]interface{}) error {
	// Check if the book exists before deleting
	_, err := biz.store.FindBook(ctx, condition)
	if err != nil {
		return err
	}

	// Delete the book using the provided condition
	if err := biz.store.DeleteBook(ctx, condition); err != nil {
		return err
	}

	return nil
}
