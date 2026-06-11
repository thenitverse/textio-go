package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	// Case 1: Valid division
	res, err := divide(100, 10)
	if err != nil {
		t.Errorf("Test Case 1 Failed: Expected nil error, got %v", err)
	}
	if res != 10 {
		t.Errorf("Test Case 1 Failed: Expected 10, got %v", res)
	}

	// Case 2: Division by zero (Triggers your custom error)
	_, err = divide(42, 0)
	if err == nil {
		t.Errorf("Test Case 2 Failed: Expected an error, got nil")
	}

	expectedError := "cannot divide 42 by zero"
	if err.Error() != expectedError {
		t.Errorf("Test Case 2 Failed: Expected message '%s', got '%s'", expectedError, err.Error())
	}
}
