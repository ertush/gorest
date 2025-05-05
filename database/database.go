package database

import (
	"github/ertush/gorest/models"
	"log"
	"os"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDB() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
