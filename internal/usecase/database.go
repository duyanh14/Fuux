package usecase

import (
	"errors"
	"fmt"
	"fuux/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func Database(config *entity.Config) (*gorm.DB, error) {
	configDatabase := config.Database

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", configDatabase.Host,
		configDatabase.User,
		configDatabase.Password,
		configDatabase.Database,
		configDatabase.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("can't connect database")
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&entity.Resource{})
	db.AutoMigrate(&entity.File{})

	return db, nil
}
