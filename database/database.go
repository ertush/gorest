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

	dsn := os.Getenv("DSN") // fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Nairobi", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	log.Println("dsn: " + dsn)

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
