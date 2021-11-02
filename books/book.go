package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/schulzf/go-tutorial/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	// db := database.DBConn
	// var book Book
	// book.Title = "My First Book"
	// book.Author = "Felipe Schulz"
	// book.Rating = 5
	// db.Create(&book)
	var data string
	c.Body()

	return c.JSON(data)
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete book")
}