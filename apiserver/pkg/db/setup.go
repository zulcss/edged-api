package db

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func InitDatabase() {
	log.Println("Initializing database...")
	database, err := gorm.Open(sqlite.Open("edged.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}

	// cluster
	database.AutoMigrate(&Cluster{})

	DB = database
}
