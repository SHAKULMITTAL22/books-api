package repository

import (
	"booksApi"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestCreate_8d57755d0c(t *testing.T) {
	// Setup
	db, err := sqlx.Open("postgres", "postgres://localhost:5432/books_test")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := &BooksPostgres{db}

	// Test case: Success
	input := booksApi.Book{
		Title:  "The Lord of the Rings",
		Author: "J.R.R. Tolkien",
		Year:   1954,
	}
	id, err := r.Create(input)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}
	if id == 0 {
		t.Error("Create() returned an invalid ID")
	}
	t.Log("Create() succeeded")

	// Test case: Failure - duplicate title
	input = booksApi.Book{
		Title:  "The Lord of the Rings",
		Author: "J.R.R. Tolkien",
		Year:   1954,
	}
	_, err = r.Create(input)
	if err == nil {
		t.Error("Create() did not return an error for a duplicate title")
	}
	t.Log("Create() failed as expected for a duplicate title")

	// Test case: Failure - invalid year
	input = booksApi.Book{
		Title:  "The Lord of the Rings",
		Author: "J.R.R. Tolkien",
		Year:   -1,
	}
	_, err = r.Create(input)
	if err == nil {
		t.Error("Create() did not return an error for an invalid year")
	}
	t.Log("Create() failed as expected for an invalid year")

	// Test case: Failure - database error
	db.Close()
	_, err = r.Create(input)
	if err == nil {
		t.Error("Create() did not return an error for a database error")
	}
	t.Log("Create() failed as expected for a database error")
}
