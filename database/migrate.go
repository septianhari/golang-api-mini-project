package database

import (
	"github.com/septianhari/golang-api-mini-project/models"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func DbMigrate(dbParam *gorm.DB) {
	DbConnection = dbParam

	DbConnection.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Store{},
		&models.Category{},
		&models.Product{},
		&models.ProductLog{},
		&models.Transaction{},
		&models.TransactionDetail{},
		&models.UserCredentials{},
	)
}
