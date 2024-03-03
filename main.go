package main

import (
	"errors"
	"fmt"
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
	{ID: "1", Title: "Wealth is not enogh", Author: "Kwame Dagadu", Quantity: 5},
	{ID: "2", Title: "The Richman of Arimathea", Author: "Yaw Fragada", Quantity: 10},
	{ID: "3", Title: "Great is our potential", Author: "Kwasi Adu", Quantity: 20},
	{ID: "4", Title: "Never judge a book by its cover", Author: "Kwaku lawson", Quantity: 30},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusOK, newBook)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	fmt.Print("entered")
	id, ok := c.GetQuery("id")

	fmt.Print("id")
	fmt.Print(id)

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing Id query parameter"})
		return
	}

	book, err := getById(id)
	fmt.Print(book)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func returnBook(c *gin.Context) {
	fmt.Print("entered")
	id, ok := c.GetQuery("id")

	fmt.Print("id")
	fmt.Print(id)

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing Id query parameter"})
		return
	}

	book, err := getById(id)
	fmt.Print(book)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
