package repository

import (
	"fuux/internal/entity"
)

type File struct {
	database *entity.Database
}

func NewFile(database *entity.Database) (*File, error) {
	return &File{
		database: database,
	}, nil
}
