package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In search of lost time", Author: "John Mart", Quantity: 6},
	{ID: "2", Title: "In Packtime", Author: "Mart", Quantity: 6},
	{ID: "1", Title: "In lost time", Author: "John lent", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id key"})

		book, err := getBookById(id)

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			return
		}
		if book.Quantity == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		}

		book.Quantity -= 1
		c.IndentedJSON(http.StatusOK, book)

	}

}
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found ")

}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	r := gin.Default()

	r.GET("/books", getBooks)

	r.POST("/books", createBook)
	r.GET("books/:id", bookById)

	r.Run("localhost:8080")

}
