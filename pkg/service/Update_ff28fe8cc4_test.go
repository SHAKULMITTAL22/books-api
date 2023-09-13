package service

import (
	"booksApi"
	"booksApi/pkg/repository"
	"testing"
)

func TestUpdate_ff28fe8cc4(t *testing.T) {
	// Create a mock repository.
	repo := &repository.MockRepository{}

	// Create a new book service.
	service := &BookService{
		repo: repo,
	}

	// Create a new book input.
	input := booksApi.UpdateBookInput{
		Title: "New Title",
		Author: "New Author",
	}

	// Call the Update method.
	err := service.Update(1, input)

	// Assert that the repository's Update method was called.
	if !repo.UpdateCalled {
		t.Error("repository's Update method was not called")
	}

	// Assert that the error is nil.
	if err != nil {
		t.Errorf("error returned from Update: %v", err)
	}

	// Assert that the book was updated.
	if repo.Book.Title != input.Title || repo.Book.Author != input.Author {
		t.Errorf("book was not updated: expected %v, got %v", input, repo.Book)
	}

	t.Log("TestUpdate_ff28fe8cc4 PASSED")
}

func TestUpdate_Error(t *testing.T) {
	// Create a mock repository.
	repo := &repository.MockRepository{}

	// Create a new book service.
	service := &BookService{
		repo: repo,
	}

	// Create a new book input.
	input := booksApi.UpdateBookInput{
		Title: "New Title",
		Author: "New Author",
	}

	// Set the repository's Update method to return an error.
	repo.UpdateError = true

	// Call the Update method.
	err := service.Update(1, input)

	// Assert that the repository's Update method was called.
	if !repo.UpdateCalled {
		t.Error("repository's Update method was not called")
	}

	// Assert that the error is not nil.
	if err == nil {
		t.Error("no error returned from Update")
	}

	// Assert that the book was not updated.
	if repo.Book.Title != input.Title || repo.Book.Author != input.Author {
		t.Errorf("book was updated: expected %v, got %v", input, repo.Book)
	}

	t.Log("TestUpdate_Error PASSED")
}
