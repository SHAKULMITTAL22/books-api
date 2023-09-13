package handler

import (
	"booksApi"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testing"
)

func TestUpdateBook_9497d5293a(t *testing.T) {
	// Test case 1: Success
	t.Log("Test case 1: Success")
	h := &Handler{
		services: &MockServices{
			Book: &MockBookService{
				UpdateFunc: func(id int, input booksApi.UpdateBookInput) error {
					return nil
				},
			},
		},
	}
	c, _ := gin.CreateTestContext(t)
	c.Request, _ = http.NewRequest(http.MethodPut, "/books/1", nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	if err := h.updateBook(c); err != nil {
		t.Errorf("Test case 1 failed: %v", err)
	} else {
		t.Log("Test case 1 passed")
	}

	// Test case 2: Invalid book id param
	t.Log("Test case 2: Invalid book id param")
	h = &Handler{
		services: &MockServices{
			Book: &MockBookService{
				UpdateFunc: func(id int, input booksApi.UpdateBookInput) error {
					return nil
				},
			},
		},
	}
	c, _ = gin.CreateTestContext(t)
	c.Request, _ = http.NewRequest(http.MethodPut, "/books/a", nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "a"}}
	if err := h.updateBook(c); err == nil {
		t.Errorf("Test case 2 failed: expected error but got nil")
	} else if err.Error() != "invalid book id param" {
		t.Errorf("Test case 2 failed: expected error 'invalid book id param' but got '%v'", err.Error())
	} else {
		t.Log("Test case 2 passed")
	}

	// Test case 3: Invalid JSON payload
	t.Log("Test case 3: Invalid JSON payload")
	h = &Handler{
		services: &MockServices{
			Book: &MockBookService{
				UpdateFunc: func(id int, input booksApi.UpdateBookInput) error {
					return nil
				},
			},
		},
	}
	c, _ = gin.CreateTestContext(t)
	c.Request, _ = http.NewRequest(http.MethodPut, "/books/1", nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte("invalid json")))
	if err := h.updateBook(c); err == nil {
		t.Errorf("Test case 3 failed: expected error but got nil")
	} else if err.Error() != "invalid JSON payload" {
		t.Errorf("Test case 3 failed: expected error 'invalid JSON payload' but got '%v'", err.Error())
	} else {
		t.Log("Test case 3 passed")
	}

	// Test case 4: Error updating book
	t.Log("Test case 4: Error updating book")
	h = &Handler{
		services: &MockServices{
			Book: &MockBookService{
				UpdateFunc: func(id int, input booksApi.UpdateBookInput) error {
					return errors.New("error updating book")
				},
			},
		},
	}
	c, _ = gin.CreateTestContext(t)
	c.Request, _ = http.NewRequest(http.MethodPut, "/books/1", nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	if err := h.updateBook(c); err == nil {
		t.Errorf("Test case 4 failed: expected error but got nil")
	} else if err.Error() != "error updating book" {
		t.Errorf("Test case 4 failed: expected error 'error updating book' but got '%v'", err.Error())
	} else {
		t.Log("Test case 4 passed")
	}
}
