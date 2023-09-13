package handler

import (
	"booksApi"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TestDeleteBook_93cbfc748f(t *testing.T) {
	// Test case 1: Success
	id := 1
	c, _ := gin.CreateTestContext(t)
	c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(id)}}
	h := &Handler{services: &booksApi.Services{Book: &booksApi.BookService{}}}
	h.deleteBook(c)
	if c.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, c.Writer.Status())
	}

	// Test case 2: Invalid book id param
	c, _ = gin.CreateTestContext(t)
	c.Params = []gin.Param{{Key: "id", Value: "invalid"}}
	h = &Handler{services: &booksApi.Services{Book: &booksApi.BookService{}}}
	h.deleteBook(c)
	if c.Writer.Status() != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, c.Writer.Status())
	}

	// Test case 3: Error deleting book
	c, _ = gin.CreateTestContext(t)
	c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(id)}}
	h = &Handler{services: &booksApi.Services{Book: &booksApi.BookService{}}}
	h.services.Book.Delete = func(id int) error {
		return errors.New("error deleting book")
	}
	h.deleteBook(c)
	if c.Writer.Status() != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, c.Writer.Status())
	}
}
