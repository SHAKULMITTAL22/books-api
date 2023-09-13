package repository

import (
	"booksApi"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestNewRepository_f43063d8cb(t *testing.T) {
	// Create a new sqlx database connection
	db, err := sqlx.Connect("postgres", "user=postgres password=mysecret dbname=booksapi sslmode=disable")
	if err != nil {
		t.Error(err)
	}

	// Create a new repository
	repo := NewRepository(db)

	// Test that the repository was created successfully
	if repo == nil {
		t.Error("Repository was not created successfully")
	}

	// Test that the repository has the correct database connection
	if repo.DB != db {
		t.Errorf("Repository has the wrong database connection: %v", repo.DB)
	}

	// Test that the repository has the correct book model
	if repo.Books == nil {
		t.Error("Repository does not have the correct book model")
	}

	// Test that the repository has the correct book model
	if repo.Book == nil {
		t.Error("Repository does not have the correct book model")
	}
}
