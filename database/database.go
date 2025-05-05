package database

import (
	"github/ertush/gorest/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// "gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDB() {

	db, err := gorm.Open(sqlite.Open("rest.db"), &gorm.Config{})

	// log.Printf("Db: %v\n", db)

	if err != nil {
		log.Fatalln("Unable to Connect to DB")
		os.Exit(2)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Successfully connected to database")

	// Migrations

	if err := db.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		log.Fatalln("Unable to run migrations!")
		os.Exit(2)
	}

	Database = DbInstance{
		Db: db,
	}
}
