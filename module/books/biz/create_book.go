package bookbiz

import (
	"context"
	bookmodel "local-app/module/books/models"
)

// CreateBookStorage defines the interface for creating a book.
type CreateBookStorage interface {
	CreateBook(ctx context.Context, data *bookmodel.Book) error
}

// createBookBiz implements the business logic for creating a new book.
type createBookBiz struct {
	store CreateBookStorage // Store for book creation
}

// NewCreateBookBiz creates a new instance of createBookBiz.
func NewCreateBookBiz(store CreateBookStorage) *createBookBiz {
	return &createBookBiz{store: store}
}

// CreateBook creates a new book after validating the input data.
func (biz *createBookBiz) CreateBook(ctx context.Context, data *bookmodel.Book) error {
	// Validate the book data
	if err := data.Validate(); err != nil {
		return err // Return validation error if data is invalid
	}

	// Call the storage to create the book in the database
	if err := biz.store.CreateBook(ctx, data); err != nil {
		return err // Return any error occurred during the creation process
	}

	return nil // Return nil indicating successful creation of the book
}
