package main

import (
	m "GIN/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := Connect()
	if err != nil {
		panic(err)
	}

	r := SetupRouter(db)
	r.Run(":8888")
}

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
