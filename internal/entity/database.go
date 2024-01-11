package entity

import "gorm.io/gorm"

type Database struct {
	Postgres *gorm.DB
	//	Redis     *redis.Client
}
