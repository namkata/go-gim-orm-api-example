package bookmodel

import (
	"errors"
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
	ID        uint       `json:"id" gorm:"column:id;"`
	Name      string     `json:"name" gorm:"column:name;"`
	URL       string     `json:"url" gorm:"column:url;"`
	Price     float32    `json:"price" gorm:"column:price;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
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

// DataPaging represents the paging information for data
type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

// ProcessPaging ensures that paging values are valid
func (p *DataPaging) ProcessPaging() {
	p.setDefaultIfInvalid()
}

// setDefaultIfInvalid sets default values for paging if invalid values are provided
func (p *DataPaging) setDefaultIfInvalid() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
}
