package handler

import (
	"booksApi/pkg/service"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestInitRoutes_e2e4f663d4(t *testing.T) {
	// Test case 1: Success
	h := &Handler{
		service: &service.MockService{},
	}
	router := h.InitRoutes()
	if router == nil {
		t.Error("router is nil")
	}
	t.Log("Test case 1 passed")

	// Test case 2: Failure
	h = &Handler{
		service: nil,
	}
	router = h.InitRoutes()
	if router != nil {
		t.Errorf("router is not nil, got %v", router)
	}
	t.Log("Test case 2 passed")
}
