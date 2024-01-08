package bookmodel

import (
	"errors"
	"local-app/common"
	"time"
)

// Error definitions
var (
	ErrBlankBookName        = errors.New("book name cannot be blank")
	ErrBookNotFound         = errors.New("book not found")
	ErrCannotUpdateFinished = errors.New("cannot update finished book")
)

// Book represents a book entity
type Book struct {
	common.BaseModel `json:",inline"`
	Name             string     `json:"name" gorm:"column:name;"`
	URL              string     `json:"url" gorm:"column:url;"`
	Price            float32    `json:"price" gorm:"column:price;"`
	CreatedAt        *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt        *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

// TableName specifies the table name for the Book struct
func (Book) TableName() string {
	return "book"
}

// Validate checks if the book name is blank
func (b Book) Validate() error {
	if b.Name == "" {
		return ErrBlankBookName
	}
	return nil
}
