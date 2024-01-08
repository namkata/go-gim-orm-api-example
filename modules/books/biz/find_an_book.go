package bookbiz

import (
	"context"
	bookmodel "local-app/modules/books/models"
)

// FindBookStorage defines the interface for finding a book in storage.
type FindBookStorage interface {
	FindBook(ctx context.Context, condition map[string]interface{}) (*bookmodel.Book, error)
}

// findBookBiz represents the business logic for finding a book.
type findBookBiz struct {
	store FindBookStorage
}

// NewFindBookBiz creates a new instance of findBookBiz.
func NewFindBookBiz(store FindBookStorage) *findBookBiz {
	return &findBookBiz{store: store}
}

// FindAnBook finds a book based on provided conditions.
func (biz *findBookBiz) FindAnBook(ctx context.Context, condition map[string]interface{}) (*bookmodel.Book, error) {
	bookData, err := biz.store.FindBook(ctx, condition)
	if err != nil {
		return nil, err
	}
	return bookData, nil
}
