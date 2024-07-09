package app

import (
	"openmusic-api/helper"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// migrate -database "mysql://root@tcp(localhost:3306)/openmusic" -path db/migrations up
func OpenConnection() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/openmusic?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	sqlDB, err := db.DB()
	helper.PanicIfError(err)

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}