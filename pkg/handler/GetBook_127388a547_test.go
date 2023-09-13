package handler

import (
    "booksApi"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "testing"
)

func TestGetBook_127388a547(t *testing.T) {
    // Test case 1: Success
    c, _ := gin.CreateTestContext(t)
    c.Request = &http.Request{URL: &url.URL{Query: url.Values{"id": {"1"}}}}
    h := &Handler{services: &booksApi.Services{Book: &booksApi.BookService{}}}
    h.getBook(c)
    if c.Writer.Status() != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, c.Writer.Status())
    }

    // Test case 2: Invalid book id param
    c, _ = gin.CreateTestContext(t)
    c.Request = &http.Request{URL: &url.URL{Query: url.Values{"id": {"a"}}}}
    h = &Handler{services: &booksApi.Services{Book: &booksApi.BookService{}}}
    h.getBook(c)
    if c.Writer.Status() != http.StatusBadRequest {
        t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, c.Writer.Status())
    }

    // Test case 3: Book not found
    c, _ = gin.CreateTestContext(t)
    c.Request = &http.Request{URL: &url.URL{Query: url.Values{"id": {"100"}}}}
    h = &Handler{services: &booksApi.Services{Book: &booksApi.BookService{}}}
    h.getBook(c)
    if c.Writer.Status() != http.StatusInternalServerError {
        t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, c.Writer.Status())
    }
}
