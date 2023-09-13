package repository

import (
	"booksApi"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestGetById_5b61de6ac5(t *testing.T) {
	// Setup
	db, err := sqlx.Open("postgres", "postgres://localhost:5432/books_test")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewBooksPostgres(db)

	// Test case: Book exists
	book, err := repo.GetById(1)
	if err != nil {
		t.Errorf("GetById(1) failed: %v", err)
	}
	if book.ID != 1 {
		t.Errorf("GetById(1) returned the wrong book: %v", book)
	}
	t.Log("GetById(1) succeeded")

	// Test case: Book does not exist
	book, err = repo.GetById(2)
	if err == nil {
		t.Errorf("GetById(2) succeeded, but should have failed")
	}
	if book.ID != 0 {
		t.Errorf("GetById(2) returned the wrong book: %v", book)
	}
	t.Log("GetById(2) failed as expected")
}
