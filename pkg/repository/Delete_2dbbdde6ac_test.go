package repository

import (
	"booksApi"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestDelete_2dbbdde6ac(t *testing.T) {
	// Setup
	db, err := sqlx.Open("postgres", "postgres://localhost:5432/books_test")
	if err != nil {
		t.Fatal(err)
	}
	r := &BooksPostgres{db: db}

	// Test case: Success
	err = r.Delete(1)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	t.Log("Test case: Success passed")

	// Test case: Failure - Record not found
	err = r.Delete(2)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	t.Log("Test case: Failure - Record not found passed")

	// Teardown
	db.Close()
}
