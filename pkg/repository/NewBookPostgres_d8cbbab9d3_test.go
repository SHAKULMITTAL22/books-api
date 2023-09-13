package repository

import (
	"booksApi"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestNewBookPostgres_d8cbbab9d3(t *testing.T) {
	// Create a new sqlx.DB instance.
	db, err := sqlx.Open("postgres", "user=postgres password=mysecretpassword dbname=books sslmode=disable")
	if err != nil {
		t.Error(err)
	}

	// Create a new BooksPostgres instance.
	booksPostgres := NewBookPostgres(db)

	// Assert that the BooksPostgres instance is not nil.
	if booksPostgres == nil {
		t.Error("BooksPostgres instance is nil")
	}

	// Assert that the BooksPostgres instance has a valid db field.
	if booksPostgres.db == nil {
		t.Error("BooksPostgres instance has a nil db field")
	}

	// Close the sqlx.DB instance.
	db.Close()
}
