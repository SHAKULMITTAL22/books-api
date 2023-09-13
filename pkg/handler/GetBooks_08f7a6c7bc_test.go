package handler

import (
	"booksApi"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TestGetBooks_08f7a6c7bc(t *testing.T) {
	// Test case 1: Success
	t.Log("Test case 1: Success")
	h := &Handler{
		services: &booksApi.Services{
			Books: &booksApi.BookService{
				Books: []booksApi.Book{
					{ID: 1, Title: "Book 1", Author: "Author 1"},
					{ID: 2, Title: "Book 2", Author: "Author 2"},
				},
			},
		},
	}
	c, _ := gin.CreateTestContext(t)
	h.getBooks(c)
	if c.Writer.Status() != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, c.Writer.Status())
	}

	// Test case 2: Error
	t.Log("Test case 2: Error")
	h = &Handler{
		services: &booksApi.Services{
			Books: &booksApi.BookService{
				Err: errors.New("error"),
			},
		},
	}
	c, _ = gin.CreateTestContext(t)
	h.getBooks(c)
	if c.Writer.Status() != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, c.Writer.Status())
	}
}
