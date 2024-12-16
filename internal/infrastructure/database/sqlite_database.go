package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func NewSQLiteDatabase(dsn string) (*gorm.DB, error) {
	var err error

	db, err = gorm.Open(
		sqlite.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)

	return db, err
}
