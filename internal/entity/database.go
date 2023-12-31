package entity

import "gorm.io/gorm"

type Database struct {
	Postgres  *gorm.DB
	SQLServer *gorm.DB
	//	Redis     *redis.Client
}
