package database

import (
	"fmt"
	"log"
	"os"

	"github.com/septianhari/golang-api-mini-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func ConnectDatabase() {
	// Ambil konfigurasi database dari environment variable
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DbConnection = db
	log.Println("Database connected successfully")

	// Panggil fungsi migrasi
	DbMigrate()
}

func DbMigrate() {
	err := DbConnection.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Store{},
		&models.Category{},
		&models.Product{},
		&models.ProductLog{},
		&models.ProductImage{},
		&models.Transaction{},
		&models.TransactionDetail{},
		&models.UserCredentials{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed")
}
