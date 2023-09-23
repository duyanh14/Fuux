package repository

import (
	"gorm.io/gorm"
)

type file struct {
	db *gorm.DB
}

var File *file

func NewFile(db *gorm.DB) (*file, error) {
	File = &file{
		db: db,
	}

	return File, nil
}
