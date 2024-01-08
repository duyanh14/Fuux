package postgres

import (
	"errors"
	"fmt"
	"fuux/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func Connect(cfg *entity.ConfigDatabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Can't connect Postgres")
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&entity.File{})
	db.AutoMigrate(&entity.ResourceAccess{})
	db.AutoMigrate(&entity.Resource{})

	return db, nil
}
