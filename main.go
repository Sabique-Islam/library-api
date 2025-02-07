package main

import (
	//"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct{
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	{ID: "4", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 1},
	{ID: "5", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 3},
	{ID: "6", Title: "Game of thrones", Author: "George R.R. Martin", Quantity: 2},
	{ID: "7", Title: "Harry Potter", Author: "J.K. Rowling", Quantity: 4},
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func getBooks(c *gin.Context) { //c: context
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}	

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil { //BindJSON: handles error sending response
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
	
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK,book)
}	

func main() {
	router := gin.Default()
	router.GET("/api/books", getBooks)
	router.GET("/api/books/:id", bookById)
	router.POST("/api/books", createBook)
	router.PATCH("/api/checkout", checkoutBook)
	router.PATCH("/api/return", returnBook)
	router.Run("localhost:8080")
	//fmt.Println(books[0].Title)
}