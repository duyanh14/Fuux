package repository

import (
	"gorm.io/gorm"
)

type file struct {
	Db *gorm.DB
}

var File *file

func NewFile(db *gorm.DB) (*file, error) {
	File = &file{
		Db: db,
	}

	return File, nil
}

func MatchRecord(field string, value interface{}, model interface{}) (bool, modelOut interface{}) {
	rs := File.Db.Where(field+" = ?", value).First(model)
	if rs.Error != nil {
		return false, model
	} else if rs.RowsAffected > 0 {
		return true, model
	}
	return false, model
}

func MatchRecordTableName(table, name string, value interface{}, model interface{}) bool {
	rs := File.Db.Table(table).Where(name+" = ?", value).First(model)
	if rs.RowsAffected > 0 {
		return true
	}
	return false
}
