package repository

import (
	"booksApi"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestGetAll_8252471115(t *testing.T) {
	// Setup
	db, err := sqlx.Open("postgres", "postgres://localhost:5432/books_test")
	if err != nil {
		t.Fatal(err)
	}
	r := &BooksPostgres{db: db}

	// Test case: Success
	books, err := r.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(books) == 0 {
		t.Error("Expected at least one book")
	}
	t.Log("Test case: Success")

	// Test case: Failure
	db.Close()
	books, err = r.GetAll()
	if err == nil {
		t.Error("Expected an error")
	}
	if len(books) > 0 {
		t.Error("Expected no books")
	}
	t.Log("Test case: Failure")
}
