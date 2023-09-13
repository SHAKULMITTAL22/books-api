package handler

import (
	"booksApi"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TestCreateBook_6e3c2a405c(t *testing.T) {
	// Create a new gin context
	c := gin.NewContext(nil, nil)

	// Create a new book
	book := booksApi.Book{
		Title:  "Test book",
		Author: "Test author",
	}

	// Bind the book to the context
	if err := c.BindJSON(&book); err != nil {
		t.Error(err)
		return
	}

	// Create a new handler
	h := &Handler{
		services: &Services{
			Book: &BookService{
				Create: func(book booksApi.Book) (int, error) {
					return 1, nil
				},
			},
		},
	}

	// Call the createBook method
	h.createBook(c)

	// Check the response status code
	if c.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, c.Writer.Status())
		return
	}

	// Check the response body
	body, err := c.Writer.Body.Bytes()
	if err != nil {
		t.Error(err)
		return
	}

	// Check the id of the created book
	id, err := strconv.Atoi(string(body))
	if err != nil {
		t.Error(err)
		return
	}

	if id != 1 {
		t.Errorf("Expected id 1, got %d", id)
		return
	}

	t.Log("TestCreateBook_6e3c2a405c PASSED")
}

func TestCreateBook_Error(t *testing.T) {
	// Create a new gin context
	c := gin.NewContext(nil, nil)

	// Create a new book
	book := booksApi.Book{
		Title:  "Test book",
		Author: "Test author",
	}

	// Bind the book to the context
	if err := c.BindJSON(&book); err != nil {
		t.Error(err)
		return
	}

	// Create a new handler
	h := &Handler{
		services: &Services{
			Book: &BookService{
				Create: func(book booksApi.Book) (int, error) {
					return 0, errors.New("error")
				},
			},
		},
	}

	// Call the createBook method
	h.createBook(c)

	// Check the response status code
	if c.Writer.Status() != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, c.Writer.Status())
		return
	}

	// Check the response body
	body, err := c.Writer.Body.Bytes()
	if err != nil {
		t.Error(err)
		return
	}

	// Check the error message
	if string(body) != "error" {
		t.Errorf("Expected error message \"error\", got %s", string(body))
		return
	}

	t.Log("TestCreateBook_Error PASSED")
}
