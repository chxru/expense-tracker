package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func routes(app *fiber.App) {

}

func main() {
	// read env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// connecting to planetscale
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to database")

	// initiating fiber app
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	log.Fatal(app.Listen(":" + os.Getenv("API_PORT")))
}
