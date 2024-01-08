package bookstorage

import "gorm.io/gorm"

type dbStorage struct {
	db *gorm.DB
}

func NewDBStorage(db *gorm.DB) *dbStorage {
	return &dbStorage{db: db}
}
