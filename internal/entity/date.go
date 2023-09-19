package entity

import (
	"gorm.io/gorm"
	"time"
)

type Date struct {
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func (m *Date) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

func (m *Date) BeforeCreate(tx *gorm.DB) error {
	if m.UpdatedAt == 0 {
		tx.Statement.SetColumn("UpdatedAt", time.Now().Unix())
	}

	tx.Statement.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}
