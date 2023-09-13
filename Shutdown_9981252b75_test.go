package booksApi

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestShutdown_9981252b75(t *testing.T) {
	// Create a new server.
	s := &Server{
		httpServer: &http.Server{
			Addr: ":8080",
		},
	}

	// Create a new context.
	ctx := context.Background()

	// Call the Shutdown method.
	err := s.Shutdown(ctx)

	// Check for errors.
	if err != nil {
		t.Errorf("Shutdown failed: %v", err)
	}

	// Log a message indicating that the test was successful.
	t.Log("Shutdown successful")
}

func TestShutdown_Error(t *testing.T) {
	// Create a new server.
	s := &Server{
		httpServer: &http.Server{
			Addr: ":8080",
		},
	}

	// Create a new context.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the Shutdown method.
	err := s.Shutdown(ctx)

	// Check for errors.
	if err == nil {
		t.Errorf("Shutdown should have failed")
	}

	// Log a message indicating that the test was successful.
	t.Log("Shutdown failed as expected")
}
