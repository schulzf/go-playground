package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	book "github.com/schulzf/go-tutorial/books"
	"github.com/schulzf/go-tutorial/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=pgsql dbname=dev port=15432 sslmode=disable TimeZone=CET"

	database.DBConn, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println(("Pgsql connected"))
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("DB Migration successful")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	app.Listen(":3000")
}