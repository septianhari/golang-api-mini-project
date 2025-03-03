package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/septianhari/golang-api-mini-project/database"
	"github.com/septianhari/golang-api-mini-project/routers"
)

var (
	DB  *gorm.DB
	err error
)

func main() {
	//DEPLOY VERSION
	//err = godotenv.Load("config/deploy.env")

	// LOCALHOST VERSION
	err = godotenv.Load("config/local.env")

	if err != nil {
		log.Fatal("ERROR loading .env file")
	} else {
		fmt.Println("Loading .env file SUCCESS")
	}

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, os.Getenv("user"), os.Getenv("password"), os.Getenv("host"), os.Getenv("port"), os.Getenv("dbname"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("ERROR connecting database")
	} else {
		fmt.Println("Connecting databse SUCCESS")
	}

	database.DbMigrate()

	app := fiber.New()

	//routing
	routers.Routing(app)

	app.Listen(":3000")
}
