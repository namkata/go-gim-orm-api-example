package userstorage

import (
	"context"

	usermodel "local-app/modules/users/models"
)

// CreateBook creates a new book entry in the database.
func (s *dbStorage) CreateUser(ctx context.Context, data *usermodel.User) error {
	// Attempt to create a new entry in the database using the provided Book data
	if err := s.db.Create(data).Error; err != nil {
		return err // Return the error if there's any issue in creating the book entry
	}
	return nil // Return nil indicating success if no error occurred
}
