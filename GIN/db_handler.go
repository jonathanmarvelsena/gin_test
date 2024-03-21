package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	m "GIN/model"
)

var db *gorm.DB

func Connect() error {
	var err error
	// Open a database connection
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto Migrate the Database
	db.AutoMigrate(&m.Item{})

	return nil
}
