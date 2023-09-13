package service

import (
	"booksApi"
	"booksApi/pkg/repository"
	"testing"
)

func TestGetById_696d38ded7(t *testing.T) {
	// Create a mock repository.
	repo := &repository.MockRepository{}

	// Create a book.
	book := booksApi.Book{
		ID:     1,
		Title:  "Harry Potter and the Sorcerer's Stone",
		Author: "J.K. Rowling",
	}

	// Configure the mock repository to return the book.
	repo.On("GetById", 1).Return(book, nil)

	// Create a book service.
	service := &BookService{
		repo: repo,
	}

	// Call the GetById method.
	result, err := service.GetById(1)

	// Assert that the book was returned.
	if result != book {
		t.Errorf("Expected book %v, got %v", book, result)
	}

	// Assert that there was no error.
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify that the mock repository was called.
	if !repo.AssertExpectations(t) {
		t.Errorf("Mock repository was not called as expected")
	}
}

func TestGetById_Error(t *testing.T) {
	// Create a mock repository.
	repo := &repository.MockRepository{}

	// Configure the mock repository to return an error.
	repo.On("GetById", 1).Return(booksApi.Book{}, errors.New("some error"))

	// Create a book service.
	service := &BookService{
		repo: repo,
	}

	// Call the GetById method.
	_, err := service.GetById(1)

	// Assert that there was an error.
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Assert that the error message is correct.
	if err.Error() != "some error" {
		t.Errorf("Expected error message 'some error', got '%v'", err.Error())
	}

	// Verify that the mock repository was called.
	if !repo.AssertExpectations(t) {
		t.Errorf("Mock repository was not called as expected")
	}
}
