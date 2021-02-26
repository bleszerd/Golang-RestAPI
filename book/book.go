package book

import (
	"github.com/bleszerd/Go-Rest-API/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

//GetBooks return all books from database
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)

	c.JSON(books)
}

//GetBook return just a book from database
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)

	c.JSON(book)
}

//NewBook create a new book on database
/*
required: {
	title: string,
	author: string,
	rating: int32
}
*/
func NewBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
	}

	db.Create(&book)

	db.Create(&book)
	c.JSON(book)
}

//DeleteBook remove a book from database
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found with given ID")
		return
	}

	db.Delete(&book)
	c.Send("Book successfully deleted")
}
