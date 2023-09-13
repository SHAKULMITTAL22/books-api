package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewErrorResponse_796c221e79(t *testing.T) {
	// Test case 1: Success
	c := &gin.Context{}
	statusCode := 400
	message := "Bad request"
	newErrorResponse(c, statusCode, message)
	if c.Writer.Status() != statusCode {
		t.Errorf("Expected status code %d, got %d", statusCode, c.Writer.Status())
	}
	if c.Writer.Body.String() != `{"message":"Bad request"}` {
		t.Errorf("Expected response body %s, got %s", `{"message":"Bad request"}`, c.Writer.Body.String())
	}
	// Test case 2: Failure
	c = &gin.Context{}
	statusCode = 500
	message = "Internal server error"
	newErrorResponse(c, statusCode, message)
	if c.Writer.Status() != statusCode {
		t.Errorf("Expected status code %d, got %d", statusCode, c.Writer.Status())
	}
	if c.Writer.Body.String() != `{"message":"Internal server error"}` {
		t.Errorf("Expected response body %s, got %s", `{"message":"Internal server error"}`, c.Writer.Body.String())
	}
}
