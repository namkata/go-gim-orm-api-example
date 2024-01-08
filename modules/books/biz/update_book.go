package bookbiz

import (
	"context"
	bookmodel "local-app/modules/books/models"
)

type UpdateBookStorage interface {
	FindBook(ctx context.Context, condition map[string]interface{}) (*bookmodel.Book, error)
	UpdateBook(ctx context.Context, condition map[string]interface{}, dataUpdate *bookmodel.Book) error
}

type updateBookBiz struct {
	store UpdateBookStorage
}

func NewUpdateBookBiz(store UpdateBookStorage) *updateBookBiz {
	return &updateBookBiz{store: store}
}

func (biz *updateBookBiz) UpdateBook(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *bookmodel.Book,
) error {
	oldBookData, err := biz.store.FindBook(ctx, condition)

	if err != nil {
		return err
	}

	if oldBookData == nil {
		return bookmodel.ErrBookNotFound
	}

	if err := biz.store.UpdateBook(ctx, condition, dataUpdate); err != nil {
		return err
	}

	return nil
}
